package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "postgres://test:test@127.0.0.1:5432/test?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Check if the connection is established
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create a new driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"test",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Apply the migration
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("No new migrations to apply")
		} else {
			log.Fatalf("Failed to apply migration: %v", err)
		}
	} else {
		log.Println("Migration applied successfully")
	}
}
