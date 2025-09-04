package handlers

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/untanky/homepage/internal/assets"
)

func writeEncodingHeader(header http.Header, encoding string) {
	if encoding == "none" {
		return
	}

	header.Set("Content-Encoding", encoding)
}

type assetHandler struct {
	manifest *assets.Manifest
} 

func NewAssetHandler() http.Handler {
	assetHandlerInstance := assetHandler{
		manifest: assets.GetManifest(),
	}

	handler := http.NewServeMux()
	handler.Handle("GET /{path...}", assetHandlerInstance)

	return handler
}

func (h assetHandler) negotiateEncoding(asset assets.Asset, request *http.Request) assets.File {
	file := asset.GetBestCandidate()

	if file.Encoding == "none" {
		return file
	}

	supportedEncoding := strings.Split(request.Header.Get("Accept-Encoding"), ", ")
	if slices.Contains(supportedEncoding, file.Encoding) {
		return file
	}

	for _, encoding := range supportedEncoding {
		file, err := asset.GetEncodedFile(encoding)
		if err == nil {
			return file
		}
	}

	return asset.GetBase()
}

func (h assetHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.PathValue("path")
	asset, err := h.manifest.GetAsset(path) 
	if errors.Is(err, assets.NotFound) {
		notFoundError(writer, request)
		return
	}

	file := h.negotiateEncoding(asset, request)

	reader, err := file.Open()
	if err != nil {
		slog.ErrorContext(request.Context(), "could not open file", slog.Any("error", err))
		internalServerError(writer, request)
		return
	}

	header := writer.Header()
	header.Set("Content-Type", asset.Mimetype)
	header.Set("Content-Length", strconv.FormatInt(file.Size, 10))
	writeEncodingHeader(header, file.Encoding)
	header.Set("Cache-Control", "public, max-age=31556952, immutable")
	
	writer.WriteHeader(http.StatusOK)
	_, err = io.Copy(writer, reader)
	if err != nil {
		slog.ErrorContext(request.Context(), "error copying file content", slog.Any("error", err))
	}
}

