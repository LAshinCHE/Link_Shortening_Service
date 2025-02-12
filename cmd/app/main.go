package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/LAshinCHE/Link_Shortening_Service/internal/grpc"
	"github.com/LAshinCHE/Link_Shortening_Service/internal/storage"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	dbTypePtr := flag.String(flagName, "imdb", flagDBDesctiption)

	flag.Parse()
	// TODO - вынести логику обработки флагов из main
	var urlStorage storage.URLStorage
	switch *dbTypePtr {
	case "imdb":
		urlStorage = storage.NewInMemoryURLStorage()
	case "postgre":
		urlStorage = storage.NewPostgraceURLDB()
	default:
		log.Fatal("Выбрана неправильная версия базы данных.\n Выберите из списка: postgre, imdb\n")
		return
	}
	// TODO Подумать лучше над принципом работы cервиса
	grpc.Run(ctx, urlStorage)
}
