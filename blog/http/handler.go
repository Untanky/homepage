package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/blog/sql"
	"github.com/untanky/homepage/internal/components"
	"github.com/untanky/homepage/internal/media"
)

type Repository interface {
	GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error)
}

type PostRepository interface {
	GetAllMetadata(ctx context.Context, blogID blog.BlogID) ([]*blog.PostMetadata, error)
	GetPost(ctx context.Context, blogID blog.BlogID, filter sql.Filter) (blog.Post, error)
}

type AssetRepository interface {
	GetAsset(ctx context.Context, id media.AssetID) (media.Asset, error)
}

func Handler(blogRepo Repository, postRepo PostRepository) http.Handler {
	ctrl := new(controller{
		blogRepo: blogRepo,
		postRepo: postRepo,
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", ctrl.renderPostList)
	mux.HandleFunc("GET /{slug}", ctrl.renderPost)

	return mux
}

type controller struct {
	blogRepo Repository
	postRepo PostRepository
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
	components.Page(
		components.WithChild(components.BlogPage(blogListData)),
		components.WithTitle(blg.Title),
		components.WithStylesheet("/assets/main.css"),
	).Render(request.Context(), writer)
}

func (c *controller) renderPost(writer http.ResponseWriter, request *http.Request) {
	blg, err := c.blogRepo.GetBlog(request.Context(), blog.BlogID(uuid.MustParse("5e1da57f-f1cb-44cc-88cd-815b996042cf")))
	if err != nil {
		slog.ErrorContext(request.Context(), "failed request", slog.Any("error", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	slug := request.PathValue("slug")

	post, err := c.postRepo.GetPost(request.Context(), blg.ID, sql.MatchSlug(slug))
	if err != nil {
		slog.ErrorContext(request.Context(), "failed request", slog.Any("error", err))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := components.PostData{
		Post: post,
	}

	writer.WriteHeader(http.StatusOK)
	components.Page(
		components.WithChild(components.PostPage(data)),
		components.WithTitle(fmt.Sprintf("%s - %s", post.Metadata().Title, blg.Title)),
		components.WithMetadata(components.NewMetadataFromPostMetadata(post.Metadata(), blg)),
		components.WithScript("/assets/share.js"),
		components.WithStylesheet("/assets/main.css"),
	).Render(request.Context(), writer)
}
