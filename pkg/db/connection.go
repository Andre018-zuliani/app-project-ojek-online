package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Connect establishes a connection to the PostgreSQL database and returns the database instance.
func Connect(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	return db
}