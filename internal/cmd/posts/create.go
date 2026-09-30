package posts

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
	"go.lukasgrimm.me/homepage/blog"
	"go.lukasgrimm.me/homepage/blog/sql"
	"go.lukasgrimm.me/homepage/internal/config"
	"go.lukasgrimm.me/homepage/internal/database"
	"go.lukasgrimm.me/homepage/internal/media"
	mediasql "go.lukasgrimm.me/homepage/internal/media/sql"
	"go.yaml.in/yaml/v3"
)

type memoryBlogPost struct {
	content  []byte
	metadata blog.PostMetadata
}

func (post *memoryBlogPost) Metadata() blog.PostMetadata {
	return post.metadata
}

func (post *memoryBlogPost) WriteTo(writer io.Writer) (int64, error) {
	n, err := writer.Write(post.content)
	return int64(n), err
}

func (*memoryBlogPost) Close() error {
	return nil
}

type metadata struct {
	Slug     string    `yaml:"slug"`
	Title    string    `yaml:"title"`
	Summary  string    `yaml:"summary"`
	BannerID uuid.UUID `yaml:"bannerID"`
	AuthorID uuid.UUID `yaml:"authorID"`
}

type CreateCommand struct {
	Content  string    `help:"the html content of the new blog post" type:"existingFile"`
	Metadata string    `help:"the metadata of the new blog posts" type:"existingFile"`
	BlogID   uuid.UUID `help:"the id of the blog to add the post to"`
}

func (cmd *CreateCommand) Run(ctx context.Context, cfg config.Config) error {
	content, err := os.ReadFile(cmd.Content)
	if err != nil {
		return fmt.Errorf("reading content: %w", err)
	}

	metadata, err := cmd.readMetadata()
	if err != nil {
		return fmt.Errorf("reading metadata: %w", err)
	}

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("generating blog post id: %w", err)
	}

	post := memoryBlogPost{
		content: content,
		metadata: blog.PostMetadata{
			BlogID:  blog.BlogID(cmd.BlogID),
			ID:      blog.PostID(id),
			Slug:    metadata.Slug,
			Title:   metadata.Title,
			Summary: metadata.Summary,
			Banner: media.Asset{
				ID: media.AssetID(metadata.BannerID),
			},
			Author: blog.Author{
				ID: blog.AuthorID(metadata.AuthorID),
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db, err := database.Setup(ctx, cfg.Database)
	if err != nil {
		return fmt.Errorf("setting up database: %w", err)
	}

	assetRepo := mediasql.NewMediaRepository(db)
	blogRepo := sql.NewBlogRepository(db, assetRepo)

	if err := blogRepo.Create(ctx, &post); err != nil {
		return fmt.Errorf("creating blog post: %w", err)
	}

	return nil
}

func (cmd *CreateCommand) readMetadata() (metadata, error) {
	file, err := os.Open(cmd.Metadata)
	if err != nil {
		return metadata{}, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	var md metadata
	if err := decoder.Decode(&md); err != nil {
		return metadata{}, err
	}

	return md, nil
}
