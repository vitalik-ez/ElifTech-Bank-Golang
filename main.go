package main

import (
	"log"

	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/handler"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/repository"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/pkg/service"
	"github.com/vitalik-ez/ElifTech-Bank-Golang/server"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
		log.Fatal("Error occured while running http server", err.Error())
	}

}
