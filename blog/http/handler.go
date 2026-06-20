package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/untanky/homepage/blog"
	"github.com/untanky/homepage/internal/components"
)

func Handler() http.Handler {
	ctrl := new(controller{})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", ctrl.renderPostList)
	mux.HandleFunc("GET /{slug}", ctrl.renderPost)

	return mux
}

type controller struct {
}

func (c *controller) renderPostList(writer http.ResponseWriter, request *http.Request) {
	blg := blog.Blog{
		ID:      blog.BlogID(uuid.New()),
		Title:   "Lukas' Blog",
		Summary: "A blog about technology I find interesting.",
	}

	posts := []blog.PostMetadata{
		{
			BlogID:  blg.ID,
			ID:      blog.PostID(uuid.New()),
			Title:   "My first blog post",
			Summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam id pulvinar sem. Vivamus consectetur, leo sed tincidunt varius, neque sapien laoreet justo, in mollis quam orci eget nisl. Phasellus urna purus, facilisis a posuere a, auctor in odio.",
		},
		{
			BlogID:  blg.ID,
			ID:      blog.PostID(uuid.New()),
			Title:   "My second blog post",
			Summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam id pulvinar sem. Vivamus consectetur, leo sed tincidunt varius, neque sapien laoreet justo, in mollis quam orci eget nisl. Phasellus urna purus, facilisis a posuere a, auctor in odio.",
		},
		{
			BlogID:  blg.ID,
			ID:      blog.PostID(uuid.New()),
			Title:   "My third blog post",
			Summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam id pulvinar sem. Vivamus consectetur, leo sed tincidunt varius, neque sapien laoreet justo, in mollis quam orci eget nisl. Phasellus urna purus, facilisis a posuere a, auctor in odio.",
		},
		{
			BlogID:  blg.ID,
			ID:      blog.PostID(uuid.New()),
			Title:   "My fourth blog post",
			Summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam id pulvinar sem. Vivamus consectetur, leo sed tincidunt varius, neque sapien laoreet justo, in mollis quam orci eget nisl. Phasellus urna purus, facilisis a posuere a, auctor in odio.",
		},
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
