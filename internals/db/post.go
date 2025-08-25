package db

import (
	"context"
	"database/sql"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/untanky/homepage/internals/blog"
)

const (
	postTable = "posts p"
)

type sqlPostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) blog.PostRepository {
	return sqlPostRepository{
		db: db,
	}
}

func (repo sqlPostRepository) GetAll(ctx context.Context) ([]*blog.Post, error) {
	query, args, _ := sq.Select(
		"p.id", "p.slug", "p.title", "a.id", "a.name",
		"a.picture_url", "p.content", "p.keywords",
		"p.banner_url", "p.created_at", "p.edited_at",
	).
		From(postTable).
		Join("authors a ON a.id = author_id").
		// Where(sq.Eq{"delete_at": nil}).
		ToSql()

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*blog.Post
	for rows.Next() {
		post := new(blog.Post)
		author := new(blog.Author)
		var id, keywords string
		var createdAt, editedAt int64
		err = rows.Scan(
			&id, &post.Slug, &post.Title, &author.Id, &author.Name,
			&author.PictureUrl, &post.Content, &keywords,
			&post.BannerUrl, &createdAt, &editedAt,
		)
		if err != nil {
			return nil, err
		}

		post.Id, _ = uuid.Parse(id)
		post.Author = author
		post.Keywords = strings.Split(keywords, ",")
		post.CreatedAt = time.Unix(createdAt, 0)
		post.EditedAt = time.Unix(editedAt, 0)

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo sqlPostRepository) GetByID(ctx context.Context, id uuid.UUID) (*blog.Post, error) {
	query, args, _ := sq.Select(
		"slug", "title", "a.id", "a.name",
		"a.picture_url", "content", "keywords",
		"banner_url", "create_at", "edited_at",
	).
		From(postTable).
		Join("authors a ON a.id author_id").
		Where(sq.And{sq.Eq{"delete_at": nil}, sq.Eq{"id": id.String()}}).
		Limit(1).
		ToSql()

	row := repo.db.QueryRowContext(ctx, query, args...)
	post := new(blog.Post)
	author := new(blog.Author)
	var keywords string
	var createdAt, editedAt int64
	err := row.Scan(
		&post.Slug, &post.Title, &author.Id, &author.Name,
		&author.PictureUrl, &post.Content, &keywords,
		&post.BannerUrl, &createdAt, &editedAt,
	)
	if err != nil {
		return nil, err
	}

	post.Id = id
	post.Author = author
	post.Keywords = strings.Split(keywords, ",")
	post.CreatedAt = time.Unix(createdAt, 0)
	post.EditedAt = time.Unix(editedAt, 0)

	return post, nil
}

func (repo sqlPostRepository) GetBySlug(ctx context.Context, slug string) (*blog.Post, error) {
	query, args, _ := sq.Select(
		"id", "slug", "title", "a.id", "a.name",
		"a.picture_url", "content", "keywords",
		"banner_url", "create_at", "edited_at",
	).
		From(postTable).
		Join("authors a ON a.id author_id").
		Where(sq.And{sq.Eq{"delete_at": nil}, sq.Eq{"slug": slug}}).
		Limit(1).
		ToSql()

	row := repo.db.QueryRowContext(ctx, query, args...)
	post := new(blog.Post)
	author := new(blog.Author)
	var id, keywords string
	var createdAt, editedAt int64
	err := row.Scan(
		&id, &post.Slug, &post.Title, &author.Id, &author.Name,
		&author.PictureUrl, &post.Content, &keywords,
		&post.BannerUrl, &createdAt, &editedAt,
	)
	if err != nil {
		return nil, err
	}

	post.Id, _ = uuid.Parse(id)
	post.Author = author
	post.Keywords = strings.Split(keywords, ",")
	post.CreatedAt = time.Unix(createdAt, 0)
	post.EditedAt = time.Unix(editedAt, 0)

	return post, nil
}

func (repo sqlPostRepository) Create(ctx context.Context, post *blog.Post) error {
	query, args, _ := sq.Insert(postTable).
		Columns("id", "slug", "title", "author_id", "content", "keywords", "banner_url", "created_at", "edited_at").
		Values(post.Id.String(), post.Slug, post.Title, post.Author.Id.String(), post.Content, strings.Join(post.Keywords, ","), post.BannerUrl, post.CreatedAt.Unix(), post.EditedAt.Unix()).
		ToSql()

	_, err := repo.db.ExecContext(ctx, query, args...)
	return err
}

func (repo sqlPostRepository) Update(ctx context.Context, post *blog.Post) error {
	query, args, _ := sq.Update(postTable).
		Where(sq.Eq{"id": post.Id.String()}).
		SetMap(map[string]any{
			"slug":       post.Slug,
			"title":      post.Title,
			"content":    post.Content,
			"keywords":   strings.Join(post.Keywords, ","),
			"banner_url": post.BannerUrl,
			"edited_at":  post.EditedAt.Unix(),
		}).
		ToSql()

	_, err := repo.db.ExecContext(ctx, query, args...)
	return err
}

func (repo sqlPostRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query, args, _ := sq.Update(postTable).
		Where(sq.Eq{"id": id.String()}).
		Set("deleted_at", time.Now().Unix()).
		ToSql()

	_, err := repo.db.ExecContext(ctx, query, args...)
	return err
}
