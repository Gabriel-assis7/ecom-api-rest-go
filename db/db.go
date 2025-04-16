package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewSqlStorage(driver, connStr string) (*sql.DB, error) {
	db, err := sql.Open(driver, connStr)
	if err != nil {
		log.Fatalf("Could not open database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	return db, nil
}
