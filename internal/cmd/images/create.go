package images

import (
	"context"
	"fmt"
	"image"
	"os"

	"github.com/google/uuid"
	"go.lukasgrimm.me/homepage/internal/config"
	"go.lukasgrimm.me/homepage/internal/database"
	"go.lukasgrimm.me/homepage/internal/media"
	"go.lukasgrimm.me/homepage/internal/media/sql"
)

type CreateCommand struct {
	Source     string   `help:"The path to the source of the image"`
	Path       string   `help:"The path under which the image will be accessible"`
	Widths     []uint   `help:"The widths of the image to produce; can be used multiple times, or a comma separated list"`
	MediaTypes []string `help:"The mediatypes to produce"`
}

func (cmd *CreateCommand) Run(runtimeContext context.Context, cfg config.Config) error {
	img, err := readImage(cmd.Source)
	if err != nil {
		return fmt.Errorf("reading image: %w", err)
	}

	mediaTypes := make([]media.ImageType, len(cmd.MediaTypes))
	for idx, mt := range cmd.MediaTypes {
		mediaTypes[idx] = media.ImageType(mt)
	}

	pipeline := media.ImagePipeline{
		Original:   img,
		MediaTypes: mediaTypes,
		Widths:     cmd.Widths,
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
		Path:   cmd.Path,
		Width:  uint(img.Bounds().Dx()),
		Height: uint(img.Bounds().Dy()),
	}

	databaseClient, err := database.Setup(runtimeContext, cfg.Database)
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}

	repo := sql.NewMediaRepository(databaseClient)

	if err := repo.Create(runtimeContext, asset, versions); err != nil {
		return fmt.Errorf("saving asset: %w", err)
	}

	return nil

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
