package main

import (
	"log"

	"github.com/gabriel-assis7/ecom-api-rest-go/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil)
	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
