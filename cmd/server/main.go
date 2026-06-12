package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/database"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run server", slog.Any("reason", err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	databaseClient, err := database.Setup(ctx, database.Config{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5432,
		Database: "postgres",
	})
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}
	defer func() {
		if err := databaseClient.Close(ctx); err != nil {
			slog.ErrorContext(ctx, "failed to close database client", slog.Any("error", err))
		}
	}()

	server := new(http.Server{
		Addr:    ":8080",
		Handler: blog.Handler(),
	})

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("serving http: %w", err)
	}

	return nil
}
