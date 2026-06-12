package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     uint16
	Database string
}

type Client = *pgx.Conn

func Setup(ctx context.Context, config Config) (Client, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	cfg, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parsing configuration: %w", err)
	}

	conn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	return Client(conn), nil
}
