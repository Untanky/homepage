package main

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/untanky/homepage/internal/assets"
	"github.com/untanky/homepage/internal/blog"
	"github.com/untanky/homepage/internal/db"
	"github.com/untanky/homepage/internal/handlers"
)

type staticHandler struct {
	prefix  string
	handler http.Handler
}

type MainHandler struct {
	handler        http.Handler
	staticHandlers []staticHandler
}

func (handler *MainHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for _, h := range handler.staticHandlers {
		if s, ok := strings.CutPrefix(request.URL.Path, h.prefix); ok {
			request.URL.Path = s

			h.handler.ServeHTTP(writer, request)
			return
		}
	}

	handler.handler.ServeHTTP(writer, request)
}

func (handler *MainHandler) RegisterStatic(httpPrefix string, osPath string) {
	fileServer := http.FileServer(http.Dir(osPath))

	handler.staticHandlers = append(handler.staticHandlers, staticHandler{
		prefix:  httpPrefix,
		handler: fileServer,
	})
}

func renderLandingPage(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	model := layoutModel{
		title: "Lukas Grimm",
	}

	data := landingPageModel{
		links: []landingPageLink{
			{label: "GitHub", href: "https://github.com/untanky"},
			{label: "LinkedIn", href: "https://linkedin.com/in/lukasgrimm"},
		},
	}

	layout(model, landingPage(data)).Render(request.Context(), writer)
}

func main() {
	database, err := sql.Open("sqlite3", "../data.db")
	if err != nil {
		panic(err)
	}

	driver, err := sqlite3.WithInstance(database, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	postRespository := db.NewPostRepository(database)
	postService := blog.NewPostService(postRespository)

	myHandler := &MainHandler{}
	handler := handlers.NewPostHandler(postService, assets.GetManifest())

	manifestHandler := handlers.NewAssetHandler()

	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("GET /{$}", renderLandingPage)
	muxHandler.Handle("/blog/", http.StripPrefix("/blog", handler))
	muxHandler.Handle("/assets/", manifestHandler)

	myHandler.handler = muxHandler
	myHandler.RegisterStatic("/static", "../static")

	server := &http.Server{
		Addr:    ":8080",
		Handler: myHandler,
	}

	println("Listening on port 8080...")
	err = server.ListenAndServe()
	panic(err)
}
