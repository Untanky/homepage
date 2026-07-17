package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/blog"
	bloghttp "github.com/untanky/homepage/blog/http"
	"github.com/untanky/homepage/blog/sql"
	"github.com/untanky/homepage/internal/database"
	mediahttp "github.com/untanky/homepage/internal/media/http"
	mediasql "github.com/untanky/homepage/internal/media/sql"
	"github.com/untanky/homepage/internal/telemetry"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run server", slog.Any("reason", err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	logger, err := telemetry.Setup(ctx, telemetry.Config{})
	if err != nil {
		return fmt.Errorf("setting up logging: %w", err)
	}

	logger.InfoContext(ctx, "starting server")

	logger.InfoContext(ctx, "setting up database")

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
			logger.ErrorContext(ctx, "failed to close database client", slog.Any("error", err))
		}
	}()

	logger.InfoContext(ctx, "setting up server")

	server := new(http.Server{
		Addr:    ":8080",
		Handler: telemetry.NewHandler(buildHandler(databaseClient), logger),
	})

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("serving http: %w", err)
	}

	return nil
}

func buildHandler(conn *pgx.Conn) http.Handler {
	repo := sql.NewBlogRepository(conn)
	mediaRepo := mediasql.NewMediaRepository(conn)

	mux := http.NewServeMux()

	assetHandler := http.FileServer(http.Dir("./tmp/web"))
	mux.Handle("/assets/{a...}", http.StripPrefix("/assets", assetHandler))
	mux.Handle("/media/{a...}", http.StripPrefix("/media", mediahttp.Handler(mediaRepo)))
	mux.Handle("/{a...}", bloghttp.Handler(blog.NewBlogRepositoryCache(repo), repo))

	return mux
}
