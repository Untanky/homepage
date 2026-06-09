package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/untanky/homepage/blog"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run server", slog.Any("reason", err))
		os.Exit(1)
	}
}

func run(context.Context) error {
	server := new(http.Server{
		Addr:    ":8080",
		Handler: blog.Handler(),
	})

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("serving http: %w", err)
	}

	return nil
}
