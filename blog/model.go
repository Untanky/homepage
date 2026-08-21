package blog

import (
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/untanky/homepage/internal/media"
)

type (
	AuthorID uuid.UUID
	BlogID   uuid.UUID
	PostID   uuid.UUID

	MediaID = media.AssetID
)

type Blog struct {
	ID        BlogID
	Banner    media.Asset
	Title     string
	Summary   string
	Authority string
}

func (blog Blog) PostURL(metadata PostMetadata) string {
	return fmt.Sprintf("%s/%s", blog.Authority, metadata.Slug)
}

type Author struct {
	ID      AuthorID
	Name    string
	Picture media.Asset
}

type PostMetadata struct {
	BlogID BlogID
	ID     PostID
	Slug   string

	Title   string
	Summary string
	Banner  media.Asset

	Author Author

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post interface {
	Metadata() PostMetadata
	io.WriterTo
	io.Closer
}
