package sql

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/untanky/homepage/blog"
)

type BlogRepository struct {
	db *pgx.Conn
}

func NewBlogRepository(conn *pgx.Conn) *BlogRepository {
	return &BlogRepository{
		db: conn,
	}
}

func (repo *BlogRepository) GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error) {
	result := repo.db.QueryRow(ctx, "SELECT id, title, summary FROM blogs WHERE id = $1 LIMIT 1", blogID)

	blg := blog.Blog{}
	err := result.Scan(&blg.ID, &blg.Title, &blg.Summary)
	if err != nil {
		return blog.Blog{}, fmt.Errorf("retrieving blog: %w", err)
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
			return nil, fmt.Errorf("scanning metadata: %w", err)
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
		return nil, fmt.Errorf("scanning blog post data: %w", err)
	}

	return &memoryPost{
		metadata: metadata,
		content:  bytes.NewBuffer(content),
	}, nil
}
