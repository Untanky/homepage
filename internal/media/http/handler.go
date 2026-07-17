package http

import (
	"context"
	"fmt"
	"mime"
	"net/http"
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
	requestPath := request.PathValue("path")
	foo := strings.SplitN(requestPath, "@", 2)
	path, rest := foo[0], foo[1]
	foo = strings.Split(rest, ".")
	rawScale, extension := foo[0], foo[1]

	mimetype := mime.TypeByExtension(fmt.Sprintf(".%s", extension))
	scale, err := strconv.ParseFloat(rawScale, 64)
	if err != nil {
		http.Error(writer, "could not parse float", http.StatusBadRequest)
		return
	}

	assetVersion, err := c.assetRepo.GetAssetVersion(request.Context(), path, scale, extension)
	if err != nil {
		http.Error(writer, "could not get asset version", http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", mimetype)
	writer.Header().Add("Content-Length", fmt.Sprintf("%d", assetVersion.Buffer.Len()))
	assetVersion.WriterTo(writer)
}
