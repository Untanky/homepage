package http

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/components"
)

type Repository interface {
	GetBlog(ctx context.Context, blogID blog.BlogID) (blog.Blog, error)
}

type PostRepository interface {
	GetAllMetadata(ctx context.Context, blogID blog.BlogID) ([]*blog.PostMetadata, error)
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
	components.BlogPage(blogListData,
		components.WithStylesheet("/assets/main.css"),
	).Render(request.Context(), writer)
}

func (c *controller) renderPost(writer http.ResponseWriter, request *http.Request) {
	post := memoryPost{
		metadata: blog.PostMetadata{
			BlogID:  blog.BlogID(uuid.New()),
			ID:      blog.PostID(uuid.New()),
			Title:   "My first blog post",
			Summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam id pulvinar sem. Vivamus consectetur, leo sed tincidunt varius, neque sapien laoreet justo, in mollis quam orci eget nisl. Phasellus urna purus, facilisis a posuere a, auctor in odio.",
		},
		buffer: bytes.NewBuffer([]byte("<h2 class=\"text-2xl\">Works</h2>")),
	}

	data := components.PostData{
		Post: &post,
	}

	writer.WriteHeader(http.StatusOK)
	components.PostPage(data,
		components.WithStylesheet("/assets/main.css"),
	).Render(request.Context(), writer)
}

type memoryPost struct {
	metadata blog.PostMetadata
	buffer   *bytes.Buffer
}

func (post *memoryPost) Metadata() blog.PostMetadata {
	return post.metadata
}

func (post *memoryPost) WriteTo(writer io.Writer) (int64, error) {
	return post.buffer.WriteTo(writer)
}

func (post *memoryPost) Close() error {
	return nil
}
