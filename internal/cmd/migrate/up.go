package migrate

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.lukasgrimm.me/homepage/internal/config"
)

type Up struct {
	MigrationPath string `help:"Path to where the migrations are stored" default:"./db/migrations" type:"existingFile"`
}

func (up *Up) Run(cfg config.Config) error {
	migrationPath := fmt.Sprintf("file://%s", up.MigrationPath)

	dbURL := fmt.Sprintf("pgx5://%s:%s@%s:%d/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	m, err := migrate.New(migrationPath, dbURL)
	if err != nil {
		return fmt.Errorf("setting up migrate: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("no changes to be applied")
			return nil
		}

		return fmt.Errorf("applying migrations: %w", err)
	}

	log.Println("migrations applied successfully")
	return nil
}
