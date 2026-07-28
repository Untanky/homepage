package main

import (
	"fmt"
	"image"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/untanky/homepage/internal/database"
	"github.com/untanky/homepage/internal/media"
	"github.com/untanky/homepage/internal/media/sql"
)

var (
	source     string
	path       string
	widths     []uint
	mediatypes []string
)

func init() {
	createCmd.Flags().StringVar(&source, "source", "", "The path to the source of the image")
	createCmd.Flags().StringVar(&path, "path", "", "The path under which the image will be accessible")
	createCmd.Flags().UintSliceVar(&widths, "width", []uint{}, "The widths of the image to produce; can be used multiple times, or a comma separated list")
	createCmd.Flags().StringSliceVar(&mediatypes, "mediatype", []string{}, "The mediatypes to produce")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an image",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		img, err := readImage(source)
		if err != nil {
			return fmt.Errorf("reading image: %w", err)
		}

		mediaTypes := make([]media.ImageType, len(mediatypes))
		for idx, mt := range mediatypes {
			mediaTypes[idx] = media.ImageType(mt)
		}

		pipeline := media.ImagePipeline{
			Original:   img,
			MediaTypes: mediaTypes,
			Widths:     widths,
		}

		versions, err := pipeline.Run()
		if err != nil {
			return fmt.Errorf("converting image: %w", err)
		}

		assetID, err := uuid.NewV7()
		if err != nil {
			return fmt.Errorf("generating uuid: %w", err)
		}

		asset := media.Asset{
			ID:     media.AssetID(assetID),
			Path:   path,
			Width:  uint(img.Bounds().Dx()),
			Height: uint(img.Bounds().Dy()),
		}

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

		if err := repo.Create(ctx, asset, versions); err != nil {
			return fmt.Errorf("saving asset: %w", err)
		}

		return nil
	},
}

func readImage(source string) (image.Image, error) {
	imageFile, err := os.Open(source)
	if err != nil {
		return nil, fmt.Errorf("opening source image: %w", err)
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, fmt.Errorf("decode source image: %w", err)
	}

	return img, nil
}
