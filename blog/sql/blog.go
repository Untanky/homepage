package sql

import (
	"bytes"
	"context"
	"fmt"

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
		SELECT b.id, b.title, b.summary, b.banner_id
		FROM blogs b
		WHERE b.id = $1
	`

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
		ID:      row.ID,
		Title:   row.Title,
		Summary: row.Summary,
		Banner:  banner,
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
		GROUP BY p.id, au.id, ass.id
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

func (repo *BlogRepository) GetPost(ctx context.Context, blogID blog.BlogID, postID blog.PostID) (blog.Post, error) {
	const getPostSQL = `
		SELECT p.id, p.blog_id, p.title, p.slug, p.summary, p.content, p.banner_id as banner_id, p.created_at, p.updated_at,
					 au.id as author_id, au.name as author_name, au.picture_id as author_picture_id
		FROM posts p
		JOIN authors au ON p.author_id = au.id
		WHERE p.blog_id = $1 and p.id = $2
		GROUP BY p.id, au.id, ass.id
		LIMIT 1
	`

	result, err := repo.db.Query(ctx, getPostSQL, blogID, postID)
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
