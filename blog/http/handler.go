package http

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/components"
	"github.com/untanky/homepage/internal/media"
)

type Repository interface {
	GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error)
}

type PostRepository interface {
	GetAllMetadata(ctx context.Context, blogID blog.BlogID) ([]*blog.PostMetadata, error)
	GetPost(ctx context.Context, blogID blog.BlogID, postID blog.PostID) (blog.Post, error)
}

type AssetRepository interface {
	GetAsset(ctx context.Context, id media.AssetID) (media.Asset, error)
}

func Handler(blogRepo Repository, postRepo PostRepository) http.Handler {
	ctrl := new(controller{
		blogRepo:  blogRepo,
		postRepo:  postRepo,
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", ctrl.renderPostList)
	mux.HandleFunc("GET /{slug}", ctrl.renderPost)

	return mux
}

type controller struct {
	blogRepo  Repository
	postRepo  PostRepository
}

func (c *controller) renderPostList(writer http.ResponseWriter, request *http.Request) {
	blg, err := c.blogRepo.GetBlog(request.Context(), blog.BlogID(uuid.MustParse("5e1da57f-f1cb-44cc-88cd-815b996042cf")))
	if err != nil {
		slog.ErrorContext(request.Context(), "failed request", slog.Any("error", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	posts, err := c.postRepo.GetAllMetadata(request.Context(), blg.ID)
	if err != nil {
		slog.ErrorContext(request.Context(), "failed request", slog.Any("error", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	blogListData := components.BlogListData{
		Blog:  blg,
		Posts: posts,
	}

	writer.WriteHeader(http.StatusOK)
	components.BlogPage(blogListData,
		components.WithStylesheet("/assets/main.css"),
	).Render(request.Context(), writer)
}

func (c *controller) renderPost(writer http.ResponseWriter, request *http.Request) {
	post, err := c.postRepo.GetPost(request.Context(), blog.BlogID(uuid.MustParse("5e1da57f-f1cb-44cc-88cd-815b996042cf")), blog.PostID(uuid.MustParse("a3264d4c-6e85-4009-9431-e12c132b37cb")))
	if err != nil {
		slog.ErrorContext(request.Context(), "failed request", slog.Any("error", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := components.PostData{
		Post: post,
	}

	writer.WriteHeader(http.StatusOK)
	components.PostPage(data,
		components.WithStylesheet("/assets/main.css"),
	).Render(request.Context(), writer)
}
