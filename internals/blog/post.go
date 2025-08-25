package blog

import (
	"context"
	"log/slog"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	Id         uuid.UUID
	Name       string
	PictureUrl string
}

type Post struct {
	Id        uuid.UUID
	Slug      string
	Title     string
	Author    *Author
	Content   string
	Keywords  []string
	BannerUrl string
	CreatedAt time.Time
	EditedAt  time.Time
}

type AuthorRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*Author, error)
}

type PostRepository interface {
	GetAll(ctx context.Context) ([]*Post, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Post, error)
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, post *Post) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type PostService struct {
	authorRepository AuthorRepository
	postRepository   PostRepository
}

func NewPostService(postRepository PostRepository) *PostService {
	postService := new(PostService)
	postService.postRepository = postRepository
	return postService
}

func (service *PostService) GetAll(ctx context.Context) ([]*Post, error) {
	slog.Info("getting all posts")
	return service.postRepository.GetAll(ctx)
}

func (service *PostService) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	slog.Info("getting post by slug", slog.String("slug", slug))
	return service.postRepository.GetBySlug(ctx, slug)
}

type CreatePostData struct {
	Title     string
	AuthorId  uuid.UUID
	Content   string
	Keywords  []string
	BannerUrl string
}

func slugify(s string) string {
	s = strings.ToLower(s)

	re := regexp.MustCompile("[^a-z0-9]+")
	s = re.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

func (service *PostService) Create(ctx context.Context, createPost CreatePostData) (*Post, error) {
	slog.Info("creating post")

	author, err := service.authorRepository.GetByID(ctx, createPost.AuthorId)
	if err != nil {
		return nil, err
	}

	post := new(Post)
	post.Slug = slugify(createPost.Title)
	post.Title = createPost.Title
	post.Author = author
	post.Content = createPost.Content
	post.Keywords = createPost.Keywords
	post.BannerUrl = createPost.BannerUrl
	post.CreatedAt = time.Now()

	err = service.postRepository.Create(ctx, post)
	return post, err
}

func (service *PostService) Update(ctx context.Context, post *Post) error {
	slog.Info("updating post", slog.String("id", post.Id.String()), slog.String("slug", post.Slug))

	post.Slug = slugify(post.Title)
	post.EditedAt = time.Now()
	return service.postRepository.Update(ctx, post)
}

func (service *PostService) DeleteByID(ctx context.Context, id uuid.UUID) error {
	slog.Info("deleting post", slog.String("id", id.String()))

	return service.postRepository.DeleteByID(ctx, id)
}
