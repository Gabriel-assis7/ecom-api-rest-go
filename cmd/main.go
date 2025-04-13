package main

import (
	"log"
	"os"

	"github.com/gabriel-assis7/ecom-api-rest-go/cmd/api"
	"github.com/gabriel-assis7/ecom-api-rest-go/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load .env file: %v", err)
	}

	dbUrl := os.Getenv("DATABASE_URL")
	driver := "postgres"

	db, err := db.NewSqlStorage(driver, dbUrl)
	if err != nil {
		log.Fatalf("Db error: %v", err)
	}

	defer db.Close()

	log.Println("Successfully connected to the database")

	server := api.NewApiServer(":8080", nil)
	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
