package blog

import (
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

type (
	AuthorID uuid.UUID
	MediaID  uuid.UUID
	BlogID   uuid.UUID
	PostID   uuid.UUID
)

type Blog struct {
	ID      BlogID
	Title   string
	Summary string
}

func (blog Blog) PostURL(metadata PostMetadata) string {
	return fmt.Sprintf("/%s", metadata.Slug)
}

type Author struct {
	ID        AuthorID
	Name      string
	PictureID MediaID
}

type PostMetadata struct {
	BlogID BlogID
	ID     PostID
	Slug   string

	Title    string
	Summary  string
	BannerID MediaID

	Author Author

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post interface {
	Metadata() PostMetadata
	io.WriterTo
	io.Closer
}
