package media

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/google/uuid"
)

type (
	AssetID        uuid.UUID
	AssetVersionID uuid.UUID
)

type Asset struct {
	ID       AssetID
	Name     string
	Width    uint
	Height   uint
	Versions []AssetVersion
}

type AssetVersion struct {
	Scale     float64
	MediaType string
	Buffer    *bytes.Buffer
}

func (asset *AssetVersion) WriterTo(writer io.Writer) (int64, error) {
	return asset.Buffer.WriteTo(writer)
}

func (asset Asset) GeneratePictureElement(ctx context.Context, writer io.Writer) error {
	return errors.New("not implemented")
}
