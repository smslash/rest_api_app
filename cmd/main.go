package main

import (
	"log"

	todo "github.com/smslash/rest_api_app/pkg"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occured while running HTTP server: %s", err.Error())
	}
}
