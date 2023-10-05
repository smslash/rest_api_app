package main

import (
	"fmt"
	"log"

	"github.com/smslash/rest_api_app/pkg/handler"
	"github.com/smslash/rest_api_app/pkg/repository"
	"github.com/smslash/rest_api_app/pkg/service"
	todo "github.com/smslash/rest_api_app/tmp"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	fmt.Println("started")

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running HTTP server: %s", err.Error())
	}
}
