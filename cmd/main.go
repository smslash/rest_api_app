package main

import (
	"log"

	"github.com/smslash/rest_api_app/pkg/handler"
	todo "github.com/smslash/rest_api_app/tmp"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running HTTP server: %s", err.Error())
	}
}
