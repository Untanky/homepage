package handlers

import (
	"log/slog"
	"net/http"

	"github.com/untanky/homepage/internal/blog"
)

type Manifest interface {
  MapEntrypoints(entrypoints ...string) []string 
}

type postHandler struct {
	service  *blog.PostService
	manifest Manifest
}

func NewPostHandler(postService *blog.PostService, manifest Manifest) http.Handler {
	postHandler := postHandler{
		service:  postService,
		manifest: manifest,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /{$}", postHandler.renderOverview)
	handler.HandleFunc("GET /{slug}", postHandler.renderPost)

	return handler
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func notFoundError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func (handler postHandler) renderOverview(w http.ResponseWriter, r *http.Request) {
	posts, err := handler.service.GetAll(r.Context())
	if err != nil {
		slog.Error("could not render overview", slog.Any("error", err))
		internalServerError(w, r)
		return
	}

	layoutData := layoutModel{
		title: "Blog | Lukas Grimm",

		stylesheets: handler.manifest.MapEntrypoints("index.css"),
		scripts:     handler.manifest.MapEntrypoints("index.ts"),
	}
	data := overviewData{
		posts: posts,
	}

	w.WriteHeader(http.StatusOK)
	layout(layoutData, overview(data)).Render(r.Context(), w)
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
