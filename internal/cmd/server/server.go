package server

import (
	"context"
	"fmt"
	"net/http"

	"go.lukasgrimm.me/homepage/blog"
	bloghttp "go.lukasgrimm.me/homepage/blog/http"
	"go.lukasgrimm.me/homepage/blog/sql"
	"go.lukasgrimm.me/homepage/internal/config"
	"go.lukasgrimm.me/homepage/internal/database"
	mediahttp "go.lukasgrimm.me/homepage/internal/media/http"
	mediasql "go.lukasgrimm.me/homepage/internal/media/sql"
	"go.lukasgrimm.me/homepage/internal/telemetry"
)

type ServeCommand struct{}

func (cmd *ServeCommand) Run(ctx context.Context, cfg config.Config) error {
	logger, err := telemetry.Setup(ctx, telemetry.Config{})
	if err != nil {
		return fmt.Errorf("setting up logging: %w", err)
	}

	logger.InfoContext(ctx, "starting server")

	logger.InfoContext(ctx, "setting up database")

	databaseClient, err := database.Setup(ctx, cfg.Database)
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}
	defer databaseClient.Close()

	logger.InfoContext(ctx, "setting up server")

	server := new(http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: telemetry.NewHandler(buildHandler(databaseClient), logger),
	})

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("serving http: %w", err)
	}

	return nil
}

func buildHandler(conn database.Client) http.Handler {
	mediaRepo := mediasql.NewMediaRepository(conn)
	repo := sql.NewBlogRepository(conn, mediaRepo)

	mux := http.NewServeMux()

	assetHandler := http.FileServer(http.Dir("./tmp/web"))
	mux.Handle("/assets/{a...}", http.StripPrefix("/assets", assetHandler))
	mux.Handle("/media/{a...}", http.StripPrefix("/media", mediahttp.Handler(mediaRepo)))
	mux.Handle("/{a...}", bloghttp.Handler(blog.NewBlogRepositoryCache(repo), repo))

	return mux
}
