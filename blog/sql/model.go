package sql

import (
	"bytes"
	"io"

	"github.com/google/uuid"
	"github.com/untanky/homepage/blog"
)

type blogRow struct {
	ID             uuid.UUID `db:"id"`
	Title          string    `db:"title"`
	Summary        string    `db:"summary"`
	BannerID       uuid.UUID `db:"banner_id"`
	BannerPath     string    `db:"banner_path"`
	BannerScale    float64   `db:"banner_scale"`
	BannerMimetype string    `db:"banner_mimetype"`
}

type memoryPost struct {
	metadata blog.PostMetadata
	content  *bytes.Buffer
}

func (post *memoryPost) Metadata() blog.PostMetadata {
	return post.metadata
}

func (post *memoryPost) WriteTo(writer io.Writer) (int64, error) {
	return post.content.WriteTo(writer)
}

func (post *memoryPost) Close() error {
	return nil
}
