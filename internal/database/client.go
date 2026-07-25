package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     uint16
	Database string
}

type Client = *pgxpool.Pool

func Setup(ctx context.Context, config Config) (Client, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parsing configuration: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("creating pool: %w", err)
	}

	return Client(pool), nil
}
