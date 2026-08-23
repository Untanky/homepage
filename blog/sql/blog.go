package sql

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/database"
	myerrors "github.com/untanky/homepage/internal/errors"
	"github.com/untanky/homepage/internal/media"
)

type AssetRepository interface {
	GetAsset(ctx context.Context, assetID media.AssetID) (media.Asset, error)
}

type BlogRepository struct {
	db        database.Client
	assetRepo AssetRepository
}

func NewBlogRepository(db database.Client, assetRepo AssetRepository) *BlogRepository {
	return &BlogRepository{
		db:        db,
		assetRepo: assetRepo,
	}
}

func (repo *BlogRepository) GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error) {
	const getBlogSQL = `
		SELECT b.id, b.title, b.summary, b.banner_id, b.authority
		FROM blogs b
		WHERE b.id = $1
	`

	slog.Info(fmt.Sprintf("%%%s\n", blogID))

	result, err := repo.db.Query(ctx, getBlogSQL, blogID)
	if err != nil {
		return blog.Blog{}, fmt.Errorf("querying blogs: %w", err)
	}

	row, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[blogRow])
	if err != nil {
		return blog.Blog{}, fmt.Errorf("reading rows: %w", err)
	}

	banner, err := repo.assetRepo.GetAsset(ctx, row.BannerID)
	if err != nil {
		return blog.Blog{}, fmt.Errorf("finding banner: %w", err)
	}

	blg := blog.Blog{
		ID:        row.ID,
		Title:     row.Title,
		Summary:   row.Summary,
		Banner:    banner,
		Authority: row.Authority,
	}

	return blg, nil
}

func (repo *BlogRepository) GetBlogByHost(ctx context.Context, host string) (blog.Blog, error) {
	const getBlogByHostSQL = `
		SELECT b.id, b.title, b.summary, b.banner_id, b.authority
		FROM blogs b
		WHERE b.authority ilike $1
	`

	slog.Info(fmt.Sprintf("%%%s\n", host))

	result, err := repo.db.Query(ctx, getBlogByHostSQL, fmt.Sprintf("%%%s", host))
	if err != nil {
		return blog.Blog{}, fmt.Errorf("querying blogs: %w", err)
	}

	row, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[blogRow])
	if err != nil {
		return blog.Blog{}, fmt.Errorf("reading rows: %w", err)
	}

	banner, err := repo.assetRepo.GetAsset(ctx, row.BannerID)
	if err != nil {
		return blog.Blog{}, fmt.Errorf("finding banner: %w", err)
	}

	blg := blog.Blog{
		ID:        row.ID,
		Title:     row.Title,
		Summary:   row.Summary,
		Banner:    banner,
		Authority: row.Authority,
	}

	return blg, nil

}

func (repo *BlogRepository) GetAllMetadata(ctx context.Context, blogID blog.BlogID) ([]*blog.PostMetadata, error) {
	const getAllMetadataSQL = `
		SELECT p.id, p.blog_id, p.title, p.slug, p.summary, p.banner_id, p.created_at, p.updated_at,
					 au.id as author_id, au.name as author_name, au.picture_id as author_picture_id
		FROM posts p
		JOIN authors au ON p.author_id = au.id
		WHERE p.blog_id = $1
		GROUP BY p.id, au.id
	`

	result, err := repo.db.Query(ctx, getAllMetadataSQL, blogID)
	if err != nil {
		return nil, fmt.Errorf("querying all metadata: %w", err)
	}

	rows, err := pgx.CollectRows(result, pgx.RowToStructByName[postMetadataRow])
	if err != nil {
		return nil, fmt.Errorf("reading rows: %w", myerrors.InternalServerError(err))
	}

	posts := make([]*blog.PostMetadata, len(rows))
	for idx, row := range rows {
		banner, err := repo.assetRepo.GetAsset(ctx, row.BannerID)
		if err != nil {
			return nil, fmt.Errorf("getting banner: %w", myerrors.InternalServerError(err))
		}

		authorPicture, err := repo.assetRepo.GetAsset(ctx, row.AuthorPictureID)
		if err != nil {
			return nil, fmt.Errorf("getting author picture: %w", myerrors.InternalServerError(err))
		}

		posts[idx] = new(blog.PostMetadata{
			BlogID: row.BlogID,
			ID:     row.ID,
			Slug:   row.Slug,

			Title:   row.Title,
			Summary: row.Summary,
			Banner:  banner,

			Author: blog.Author{
				ID:      row.AuthorID,
				Name:    row.AuthorName,
				Picture: authorPicture,
			},

			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return posts, nil
}

func (repo *BlogRepository) GetPost(ctx context.Context, blogID blog.BlogID, filter Filter) (blog.Post, error) {
	const getPostSQL = `
		SELECT p.id, p.blog_id, p.title, p.slug, p.summary, p.content, p.banner_id as banner_id, p.created_at, p.updated_at,
					 au.id as author_id, au.name as author_name, au.picture_id as author_picture_id
		FROM posts p
		JOIN authors au ON p.author_id = au.id
		WHERE p.blog_id = $1 and %s
		GROUP BY p.id, au.id
		LIMIT 1
	`

	args := append([]any{blogID}, filter.Args()...)

	result, err := repo.db.Query(ctx, fmt.Sprintf(getPostSQL, filter.Condition()), args...)
	if err != nil {
		return nil, fmt.Errorf("querying posts: %w", handleErr(err))
	}

	row, err := pgx.CollectExactlyOneRow(result, pgx.RowToStructByName[postRow])
	if err != nil {
		return nil, fmt.Errorf("reading rows: %w", myerrors.InternalServerError(err))
	}

	banner, err := repo.assetRepo.GetAsset(ctx, row.BannerID)
	if err != nil {
		return nil, fmt.Errorf("getting banner: %w", myerrors.InternalServerError(err))
	}

	authorPicture, err := repo.assetRepo.GetAsset(ctx, row.AuthorPictureID)
	if err != nil {
		return nil, fmt.Errorf("getting author picture: %w", myerrors.InternalServerError(err))
	}

	metadata := blog.PostMetadata{
		BlogID: row.BlogID,
		ID:     row.ID,
		Slug:   row.Slug,

		Title:   row.Title,
		Summary: row.Summary,
		Banner:  banner,

		Author: blog.Author{
			ID:      row.AuthorID,
			Name:    row.AuthorName,
			Picture: authorPicture,
		},

		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return &memoryPost{
		metadata: metadata,
		content:  bytes.NewBufferString(row.Content),
	}, nil
}

func (repo *BlogRepository) Create(ctx context.Context, post blog.Post) error {
	const createBlogPost = `
		INSERT INTO posts (id, blog_id, author_id, slug, title, summary, content, banner_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	metadata := post.Metadata()

	buffer := bytes.NewBuffer(nil)
	if _, err := post.WriteTo(buffer); err != nil {
		return fmt.Errorf("reading post: %w", err)
	}

	if _, err := repo.db.Exec(ctx, createBlogPost,
		metadata.ID,
		metadata.BlogID,
		metadata.Author.ID,
		metadata.Slug,
		metadata.Title,
		metadata.Summary,
		buffer.String(),
		metadata.Banner.ID,
		metadata.CreatedAt,
		metadata.UpdatedAt,
	); err != nil {
		return fmt.Errorf("inserting blog post: %w", err)
	}

	return nil
}
