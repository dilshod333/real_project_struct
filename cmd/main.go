package main

import (
	"gin/api"
	"gin/storage"
	"gin/storage/postgres"
	"log"
)

func main() {

	db := postgres.DB()

	storage := storage.NewStoragePg(db)

	server := api.New(api.Option{
		Storage: storage,
	})

	if err := server.Run(":8080"); err != nil {
		log.Fatal("Failed to run HTTP server:  ", err)
		panic(err)
	}

	log.Print("Server stopped")

}
