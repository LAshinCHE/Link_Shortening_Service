package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/grpc"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	godotenv.Load()
	port := os.Getenv("PORT")
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
	// TODO Подумать лучше над принципом работы cервиса
	grpc.Run(ctx, urlStorage)
}
