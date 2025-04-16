package main

import (
	"log"
	"os"

	"github.com/gabriel-assis7/ecom-api-rest-go/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	db, err := db.NewSqlStorage("postgres", "postgres://postgres:9009@localhost:5432/ecom_golang?sslmode=disable")
	if err != nil {
		log.Fatalf("Db error: %v", err)
	}

	defer db.Close()

	driverInstance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driverInstance,
	)
	if err != nil {
		log.Fatalf("Migration setup failed: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide an argument: up or down")
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	log.Println("Successfully applied the operation")
}
