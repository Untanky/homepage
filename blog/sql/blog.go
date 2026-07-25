package sql

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/blog"
	myerrors "github.com/untanky/homepage/internal/errors"
	"github.com/untanky/homepage/internal/media"
)

var postNotFound = errors.New("blog post not found")

type BlogRepository struct {
	db *pgx.Conn
}

func NewBlogRepository(conn *pgx.Conn) *BlogRepository {
	return &BlogRepository{
		db: conn,
	}
}

func (repo *BlogRepository) GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error) {
	const blogQuery = `
		SELECT b.id, b.title, b.summary, b.banner_id, a.name as banner_path, av.scale as banner_scale, av.mimetype as banner_mimetype
		FROM blogs b
		JOIN media.assets a ON b.banner_id = a.id
		JOIN media.asset_versions av ON b.banner_id = av.asset_id 
		WHERE b.id = $1
	`

	result, err := repo.db.Query(ctx, blogQuery, blogID)
	if err != nil {
		return blog.Blog{}, fmt.Errorf("querying database: %w", handleErr(err))
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
	result, err := repo.db.Query(ctx, "SELECT posts.id, posts.blog_id, posts.title, posts.slug, posts.summary, posts.banner_id, posts.created_at, posts.updated_at, authors.id, authors.name, authors.picture_id FROM posts JOIN authors ON posts.author_id = authors.id WHERE posts.blog_id = $1", blogID)
	if err != nil {
		return nil, fmt.Errorf("querying all metadata: %w", err)
	}
	defer result.Close()

	metadataList := make([]*blog.PostMetadata, 0, 16)

	for result.Next() {
		post := new(blog.PostMetadata{})

		err := result.Scan(&post.ID, &post.BlogID, &post.Title, &post.Slug, &post.Summary, &post.BannerID, &post.CreatedAt, &post.UpdatedAt, &post.Author.ID, &post.Author.Name, &post.Author.PictureID)
		if err != nil {
			return nil, myerrors.InternalServerError(fmt.Errorf("scanning metadata: %w", err))
		}

		metadataList = append(metadataList, post)
	}

	return metadataList, nil
}

func (repo *BlogRepository) GetPost(ctx context.Context, blogID blog.BlogID, postID blog.PostID) (blog.Post, error) {
	row := repo.db.QueryRow(ctx, "SELECT posts.id, posts.blog_id, posts.title, posts.slug, posts.summary, posts.banner_id, posts.created_at, posts.updated_at, posts.content, authors.id, authors.name, authors.picture_id FROM posts JOIN authors on posts.author_id = authors.id WHERE posts.blog_id = $1 AND posts.id = $2", blogID, postID)

	metadata := blog.PostMetadata{}
	content := make([]byte, 0)

	err := row.Scan(&metadata.ID, &metadata.BlogID, &metadata.Title, &metadata.Slug, &metadata.Summary, &metadata.BannerID, &metadata.CreatedAt, &metadata.UpdatedAt, &content, &metadata.Author.ID, &metadata.Author.Name, &metadata.Author.PictureID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, myerrors.NotFoundError(fmt.Errorf("blog id: %s, post id: %s: %w", blogID, postID, postNotFound))
		}

		return nil, myerrors.InternalServerError(fmt.Errorf("scanning blog post data: %w", err))
	}

	return &memoryPost{
		metadata: metadata,
		content:  bytes.NewBuffer(content),
	}, nil
}
