package handlers

import (
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/untanky/homepage/internals/file"
)

func NewManifestHandler(manifest *file.Manifest) http.Handler {
	controller := manifestController{
		manifest: manifest,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /{path...}", controller.getFile)

	return handler
}

type manifestController struct {
	manifest *file.Manifest
}

func (mc manifestController) getFile(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue("path")
	file, err := mc.manifest.GetFile(path)
	if err != nil {
		slog.Error("could not find file", slog.String("path", path), slog.Any("error", err))
		notFoundError(w, r)
		return
	}

	lastModifiedString := r.Header.Get("If-Modified-Since")
	lastModifiedTime, err := time.Parse(time.RFC1123, lastModifiedString)
	if err == nil && file.LastModified.Truncate(time.Second).Equal(lastModifiedTime.UTC()) {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	reader, err := file.Open()
	if err != nil {
		slog.Error("could not get file", slog.String("path", path), slog.Any("error", err))
		internalServerError(w, r)
		return
	}
	defer reader.Close()

	header := w.Header()
	header.Add("Content-Length", strconv.FormatInt(file.Size, 10))
	header.Add("Content-Type", file.MimeType)
	header.Add("Last-Modified", file.LastModified.Format(time.RFC1123))
	w.WriteHeader(http.StatusOK)
	_, err = io.Copy(w, reader)
	if err != nil {
		slog.ErrorContext(r.Context(), "could not send file")
	}
}
