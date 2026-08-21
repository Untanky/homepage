package sql

import (
	"bytes"
	"io"
	"time"

	"github.com/untanky/homepage/blog"
)

type blogRow struct {
	ID        blog.BlogID  `db:"id"`
	Title     string       `db:"title"`
	Summary   string       `db:"summary"`
	BannerID  blog.MediaID `db:"banner_id"`
	Authority string       `db:"authority"`
}

type postMetadataRow struct {
	ID        blog.PostID  `db:"id"`
	BlogID    blog.BlogID  `db:"blog_id"`
	Slug      string       `db:"slug"`
	Title     string       `db:"title"`
	Summary   string       `db:"summary"`
	BannerID  blog.MediaID `db:"banner_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`

	AuthorID        blog.AuthorID `db:"author_id"`
	AuthorName      string        `db:"author_name"`
	AuthorPictureID blog.MediaID  `db:"author_picture_id"`
}

type postRow struct {
	ID        blog.PostID  `db:"id"`
	BlogID    blog.BlogID  `db:"blog_id"`
	Slug      string       `db:"slug"`
	Title     string       `db:"title"`
	Summary   string       `db:"summary"`
	Content   string       `db:"content"`
	BannerID  blog.MediaID `db:"banner_id"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`

	AuthorID        blog.AuthorID `db:"author_id"`
	AuthorName      string        `db:"author_name"`
	AuthorPictureID blog.MediaID  `db:"author_picture_id"`
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
