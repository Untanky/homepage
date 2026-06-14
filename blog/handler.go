package blog

import (
	"net/http"

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
	writer.WriteHeader(http.StatusOK)
	components.HTML(components.HTMLData{
		Title: "Blog",

		Scripts: []string{"https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"},
	}).Render(request.Context(), writer)
}

func (c *controller) renderPost(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("post"))
}
