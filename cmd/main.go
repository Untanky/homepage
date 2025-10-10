package main

import (
	"database/sql"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/untanky/homepage/internal/assets"
	"github.com/untanky/homepage/internal/blog"
	"github.com/untanky/homepage/internal/db"
	"github.com/untanky/homepage/internal/handlers"
)

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

	landingPageController := handlers.LandingPageController{
		Manifest: assets.GetManifest(),
	}

	postRespository := db.NewPostRepository(database)
	postService := blog.NewPostService(postRespository)

	handler := handlers.NewPostHandler(postService, assets.GetManifest())

	manifestHandler := handlers.NewAssetHandler()

	muxHandler := http.NewServeMux()
	muxHandler.Handle("GET /{$}", landingPageController)
	muxHandler.Handle("/blog/", http.StripPrefix("/blog", handler))
	muxHandler.Handle("/assets/", manifestHandler)
	muxHandler.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../static"))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: muxHandler,
	}

	println("Listening on port 8080...")
	err = server.ListenAndServe()
	panic(err)
}
