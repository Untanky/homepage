package main

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/alecthomas/kong"
	"github.com/google/uuid"
	"github.com/untanky/homepage/internal/config"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media/sql"
)

type listCommand struct {
	Filter string `arg:"" optional:"" help:"filter the listed image by path prefix"`
}

func (cmd *listCommand) Run(ctx *kong.Context, runtimeContext context.Context, cfg config.Config) error {
	databaseClient, err := database.Setup(runtimeContext, cfg.Database)
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}

	repo := sql.NewMediaRepository(databaseClient)

	assets, err := repo.ListAssets(runtimeContext, cmd.Filter)
	if err != nil {
		return fmt.Errorf("list images: %w", err)
	}

	w := tabwriter.NewWriter(ctx.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tPATH\tALT\tVERSIONS")
	for _, asset := range assets {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\n", uuid.UUID(asset.ID), asset.Path, asset.Alt, len(asset.Versions))
	}

	w.Flush()

	return nil

}
