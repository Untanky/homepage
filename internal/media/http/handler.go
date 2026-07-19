package http

import (
	"context"
	"fmt"
	"log"
	"mime"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/untanky/homepage/internal/media"
)

type AssetRepository interface {
	GetAssetVersion(ctx context.Context, path string, scale float64, mimetype string) (media.AssetVersion, error)
}

type controller struct {
	assetRepo AssetRepository
}

func Handler(assetRepo AssetRepository) http.Handler {
	ctrl := new(controller{
		assetRepo: assetRepo,
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{path...}", ctrl.serveAssetVersion)

	return mux
}

func (c *controller) serveAssetVersion(writer http.ResponseWriter, request *http.Request) {
	requestPath := request.URL.Path
	extension := path.Ext(requestPath)
	if extension == "" {
		http.Error(writer, "no externsion found", http.StatusBadRequest)
		return
	}

	requestPath, _ = strings.CutSuffix(requestPath, extension)

	parts := strings.Split(requestPath, "@")
	if len(parts) != 2 {
		http.Error(writer, "more than one quality marker found", http.StatusBadRequest)
		return
	}
	name, rawScale := parts[0], parts[1]

	mimetype := mime.TypeByExtension(extension)
	scale, err := strconv.ParseFloat(rawScale, 64)
	if err != nil {
		http.Error(writer, "could not parse float", http.StatusBadRequest)
		return
	}

	fmt.Println(strings.TrimPrefix(name, "/"), scale, mimetype)

	assetVersion, err := c.assetRepo.GetAssetVersion(request.Context(), strings.TrimPrefix(name, "/"), scale, mimetype)
	if err != nil {
		http.Error(writer, "could not get asset version", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writer.Header().Add("Content-Type", mimetype)
	writer.Header().Add("Content-Length", fmt.Sprintf("%d", assetVersion.Buffer.Len()))
	assetVersion.WriterTo(writer)
}
