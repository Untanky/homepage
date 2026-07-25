package sql

import (
	"bytes"
	"io"
	"time"

	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/media"
)

type blogRow struct {
	ID             blog.BlogID  `db:"id"`
	Title          string       `db:"title"`
	Summary        string       `db:"summary"`
	BannerID       blog.MediaID `db:"banner_id"`
	BannerPath     string       `db:"banner_path"`
	BannerScale    float64      `db:"banner_scale"`
	BannerMimetype string       `db:"banner_mimetype"`
}

type postMetadataRow struct {
	ID         blog.PostID   `db:"id"`
	BlogID     blog.BlogID   `db:"blog_id"`
	Slug       string        `db:"slug"`
	Title      string        `db:"title"`
	Summary    string        `db:"summary"`
	BannerID   blog.MediaID  `db:"banner_id"`
	BannerPath string        `db:"banner_path"`
	CreatedAt  time.Time     `db:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at"`
	AuthorID   blog.AuthorID `db:"author_id"`
	AuthorName string        `db:"author_name"`

	BannerVersions []media.AssetVersion `db:"banner_versions"`
}

type postRow struct {
	ID         blog.PostID   `db:"id"`
	BlogID     blog.BlogID   `db:"blog_id"`
	Slug       string        `db:"slug"`
	Title      string        `db:"title"`
	Summary    string        `db:"summary"`
	Content    string        `db:"content"`
	BannerID   blog.MediaID  `db:"banner_id"`
	BannerPath string        `db:"banner_path"`
	CreatedAt  time.Time     `db:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at"`
	AuthorID   blog.AuthorID `db:"author_id"`
	AuthorName string        `db:"author_name"`

	BannerVersions []media.AssetVersion `db:"banner_versions"`
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
