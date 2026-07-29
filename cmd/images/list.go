package main

import (
	"fmt"
	"text/tabwriter"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media/sql"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all images",
	Args:  cobra.MaximumNArgs(1),
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

		filter := ""
		if len(args) == 1 {
			filter = args[0]
		}

		assets, err := repo.ListAssets(ctx, filter)
		if err != nil {
			return fmt.Errorf("list images: %w", err)
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tPATH\tALT\tVERSIONS")
		for _, asset := range assets {
			fmt.Fprintf(w, "%s\t%s\t%s\t%d\n", uuid.UUID(asset.ID), asset.Path, asset.Alt, len(asset.Versions))
		}

		w.Flush()

		return nil
	},
}
