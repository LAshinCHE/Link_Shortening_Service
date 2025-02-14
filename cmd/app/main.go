package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/domain"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/grpc"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/handels"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/repository"
	"github.com/joho/godotenv"

	pb "github.com/LAshinCHE/Link_Shortening_Service/api/proto/shortener"
)

var (
	flagName          = "dbtype"
	flagDBDesctiption = "Выберите тип базы данных: imdb или postgrace\n"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	godotenv.Load()
	grpcPort := os.Getenv("GRPC_PORT")
	dbUrl := os.Getenv("DBURL")

	dbTypePtr := flag.String(flagName, "imdb", flagDBDesctiption)

	flag.Parse()
	// TODO - вынести логику обработки флагов из main
	var urlStorage repository.URLStorage
	var err error
	switch *dbTypePtr {
	case "imdb":
		urlStorage = repository.NewRepositoryIMDB()
	case "postgre":
		urlStorage, err = repository.NewRepositoryPg(ctx, dbUrl)
		if err != nil {
			log.Fatalf("При создание базы произошла ошибка: %v", err)
		}
	default:
		log.Fatal("Выбрана неправильная версия базы данных.\n Выберите из списка: postgre, imdb\n")
		return
	}

	shortenerServ := domain.NewShortenerService(urlStorage)
	handler := handels.NewHandler(shortenerServ)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShortenerServer(s, handler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
