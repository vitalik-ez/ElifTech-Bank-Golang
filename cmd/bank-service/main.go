package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/handler"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/repository"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/service"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while running http server", err.Error())
	}

}
