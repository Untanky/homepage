package main

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New("file://./db/migrations", "pgx5://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("failed to setup migrate: %v", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no changes to be applied")
			return
		}
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("migrations applied successfully")
}
