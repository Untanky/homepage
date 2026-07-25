package sql

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/database"
	myerrors "github.com/untanky/homepage/internal/errors"
	"github.com/untanky/homepage/internal/media"
)

type BlogRepository struct {
	db database.Client
}

func NewBlogRepository(db database.Client) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

func (repo *BlogRepository) GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error) {
	const getBlogSQL = `
		SELECT b.id, b.title, b.summary, b.banner_id, a.name as banner_path, av.scale as banner_scale, av.mimetype as banner_mimetype
		FROM blogs b
		JOIN media.assets a ON b.banner_id = a.id
		JOIN media.asset_versions av ON b.banner_id = av.asset_id 
		WHERE b.id = $1
	`

	result, err := repo.db.Query(ctx, getBlogSQL, blogID)
	if err != nil {
		return blog.Blog{}, fmt.Errorf("querying blogs: %w", handleErr(err))
	}

	rows, err := pgx.CollectRows(result, pgx.RowToStructByName[blogRow])
	if err != nil {
		return blog.Blog{}, fmt.Errorf("reading rows: %w", myerrors.InternalServerError(err))
	}

	blg := blog.Blog{}
	for idx, row := range rows {
		if idx == 0 {
			blg.ID = blog.BlogID(row.ID)
			blg.Title = row.Title
			blg.Summary = row.Summary
			blg.Banner.ID = blog.MediaID(row.BannerID)
			blg.Banner.Name = row.BannerPath
		}

		blg.Banner.Versions = append(blg.Banner.Versions, media.AssetVersion{
			Scale:     row.BannerScale,
			MediaType: row.BannerMimetype,
		})
	}

	return blg, nil
}

func (repo *BlogRepository) GetAllMetadata(ctx context.Context, blogID blog.BlogID) ([]*blog.PostMetadata, error) {
	const getAllMetadataSQL = `
		SELECT p.id, p.blog_id, p.title, p.slug, p.summary, ass.id as banner_id, ass.name as banner_path, p.created_at, p.updated_at,
					 au.id as author_id, au.name as author_name, json_agg(json_build_object('Scale', av.scale, 'Mediatype', av.mimetype)) as banner_versions
		FROM posts p
		JOIN authors au ON p.author_id = au.id
		JOIN media.assets ass ON p.banner_id = ass.id
		JOIN media.asset_versions av ON p.banner_id = av.asset_id
		WHERE p.blog_id = $1
		GROUP BY p.id, au.id, ass.id
	`

	result, err := repo.db.Query(ctx, getAllMetadataSQL, blogID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []*blog.PostMetadata{}, nil
		}

		return nil, fmt.Errorf("querying all metadata: %w", handleErr(err))
	}

	rows, err := pgx.CollectRows(result, pgx.RowToStructByName[postMetadataRow])
	if err != nil {
		return nil, fmt.Errorf("reading rows: %w", myerrors.InternalServerError(err))
	}

	posts := make([]*blog.PostMetadata, len(rows))
	for idx, row := range rows {
		posts[idx] = new(blog.PostMetadata{
			BlogID: row.BlogID,
			ID:     row.ID,
			Slug:   row.Slug,

			Title:   row.Title,
			Summary: row.Summary,
			Banner: media.Asset{
				ID:       row.BannerID,
				Name:     row.BannerPath,
				Versions: row.BannerVersions,
			},

			Author: blog.Author{
				ID:   row.AuthorID,
				Name: row.AuthorName,
			},

			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}

	return posts, nil
}

func (repo *BlogRepository) GetPost(ctx context.Context, blogID blog.BlogID, postID blog.PostID) (blog.Post, error) {
	const getPostSQL = `
		SELECT p.id, p.blog_id, p.title, p.slug, p.summary, p.content, ass.id as banner_id, ass.name as banner_path, p.created_at, p.updated_at,
					 au.id as author_id, au.name as author_name,
					 json_agg(json_build_object('Scale', av.scale, 'Mediatype', av.mimetype)) as banner_versions
		FROM posts p
		JOIN authors au ON p.author_id = au.id
		JOIN media.assets ass ON p.banner_id = ass.id
		JOIN media.asset_versions av ON p.banner_id = av.asset_id
		WHERE p.blog_id = $1 and p.id = $2
		GROUP BY p.id, au.id, ass.id
		LIMIT 1
	`

	result, err := repo.db.Query(ctx, getPostSQL, blogID, postID)
	if err != nil {
		return nil, fmt.Errorf("querying posts: %w", handleErr(err))
	}

	row, err := pgx.CollectOneRow(result, pgx.RowToStructByName[postRow])
	if err != nil {
		return nil, fmt.Errorf("reading rows: %w", myerrors.InternalServerError(err))
	}

	metadata := blog.PostMetadata{
		BlogID: row.BlogID,
		ID:     row.ID,
		Slug:   row.Slug,

		Title:   row.Title,
		Summary: row.Summary,
		Banner: media.Asset{
			ID:       row.BannerID,
			Name:     row.BannerPath,
			Versions: row.BannerVersions,
		},

		Author: blog.Author{
			ID:   row.AuthorID,
			Name: row.AuthorName,
		},

		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return &memoryPost{
		metadata: metadata,
		content:  bytes.NewBufferString(row.Content),
	}, nil
}
