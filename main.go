package main

import (
	"log"

	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/handler"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/repository"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/service"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/server"
)

func main() {
	/*
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
		port := os.Getenv("PORT")*/

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while running http server", err.Error())
	}

}
