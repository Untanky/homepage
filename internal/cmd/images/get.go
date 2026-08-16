package images

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/alecthomas/kong"
	"github.com/google/uuid"
	"github.com/untanky/homepage/internal/components"
	"github.com/untanky/homepage/internal/config"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media"
	"github.com/untanky/homepage/internal/media/sql"
)

type GetCommand struct {
	ID   string `arg:"" help:"The id of the image"`
	Html bool   `help:"Output the image as HTML"`
}

func (cmd *GetCommand) Run(ctx *kong.Context, runtimeContext context.Context, cfg config.Config) error {
	databaseClient, err := database.Setup(runtimeContext, cfg.Database)
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}

	repo := sql.NewMediaRepository(databaseClient)

	filter, err := uuid.Parse(cmd.ID)
	if err != nil {
		return err
	}

	asset, err := repo.GetAsset(runtimeContext, media.AssetID(filter))
	if err != nil {
		return fmt.Errorf("list images: %w", err)
	}

	if cmd.Html {
		return components.Image(asset, 10).Render(runtimeContext, ctx.Stdout)
	}

	fmt.Fprintf(ctx.Stdout, "ID: %s'\n", uuid.UUID(asset.ID))
	fmt.Fprintf(ctx.Stdout, "Path: %s\n", asset.Path)
	fmt.Fprintf(ctx.Stdout, "Alt: %s\n", asset.Alt)
	fmt.Fprintf(ctx.Stdout, "Versions:\n")
	w := tabwriter.NewWriter(ctx.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "MIMETYPE\tDIMENSIONS")
	for _, version := range asset.Versions {
		fmt.Fprintf(w, "%s\t%dx%d\n", version.Mimetype, version.Width, version.Height)
	}

	w.Flush()

	return nil

}
