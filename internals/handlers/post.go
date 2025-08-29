package handlers

import (
	"log/slog"
	"net/http"

	"github.com/untanky/homepage/internals/blog"
)

type postHandler struct {
	service *blog.PostService
}

func NewPostHandler(postService *blog.PostService) http.Handler {
	postHandler := postHandler{
		service: postService,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /{$}", postHandler.renderOverview)
	handler.HandleFunc("GET /{slug}", postHandler.renderPost)

	return handler
}

func internalServerError(w http.ResponseWriter, r *http.Request) {

}

func (handler postHandler) renderOverview(w http.ResponseWriter, r *http.Request) {
	posts, err := handler.service.GetAll(r.Context())
	if err != nil {
		slog.Error("could not render overview", slog.Any("error", err))
		internalServerError(w, r)
		return
	}

	data := overviewData{
		posts: posts,
	}

	w.WriteHeader(http.StatusOK)
	layout(layoutModel{title: "Blog | Lukas Grimm"}, overview(data)).Render(r.Context(), w)
}

func (handler postHandler) renderPost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		slog.Error("slug is missing", slog.String("path", r.URL.Path))
		internalServerError(w, r)
		return
	}

	pst, err := handler.service.GetBySlug(r.Context(), slug)
	if err != nil {
		slog.Error("could not render post", slog.Any("error", err))
		internalServerError(w, r)
		return
	}

	data := postData{
		post: pst,
	}

	w.WriteHeader(http.StatusOK)
	layout(layoutModel{title: pst.Title}, post(data)).Render(r.Context(), w)
}
