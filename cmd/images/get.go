package main

import (
	"fmt"
	"text/tabwriter"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/untanky/homepage/internal/components"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media"
	"github.com/untanky/homepage/internal/media/sql"
)

var (
	html bool
)

func init() {
	getCmd.Flags().BoolVar(&html, "html", false, "Output the image as HTML")

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an images",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

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

		repo := sql.NewMediaRepository(databaseClient)

		filter, err := uuid.Parse(args[0])
		if err != nil {
			return err
		}

		asset, err := repo.GetAsset(ctx, media.AssetID(filter))
		if err != nil {
			return fmt.Errorf("list images: %w", err)
		}

		if html {
			return components.Image(asset, 10).Render(ctx, cmd.OutOrStdout())
		}

		fmt.Fprintf(cmd.OutOrStdout(), "ID: %s'\n", uuid.UUID(asset.ID))
		fmt.Fprintf(cmd.OutOrStdout(), "Path: %s\n", asset.Path)
		fmt.Fprintf(cmd.OutOrStdout(), "Alt: %s\n", asset.Alt)
		fmt.Fprintf(cmd.OutOrStdout(), "Versions:\n")
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "MIMETYPE\tDIMENSIONS")
		for _, version := range asset.Versions {
			fmt.Fprintf(w, "%s\t%dx%d\n", version.Mimetype, version.Width, version.Height)
		}

		w.Flush()

		return nil
	},
}
