package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"

	"github.com/untanky/homepage/internal/media"
)

type AssetRepository interface {
	GetVersion(ctx context.Context, path string, mimetype string, width uint) (media.Version, error)
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
	trimmed := strings.TrimPrefix(requestPath, "/")
	dotIdx := strings.LastIndex(trimmed, ".")
	dashIdx := strings.LastIndex(trimmed, "-")
	if dotIdx < 0 || dashIdx < 0 || dashIdx > dotIdx {
		// invalid format
	}
	path := trimmed[:dashIdx]
	ext := trimmed[dotIdx:]
	width, err := strconv.ParseUint(trimmed[dashIdx+1:dotIdx], 10, 0)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	mimetype := mime.TypeByExtension(ext)

	mediaVersion, err := c.assetRepo.GetVersion(request.Context(), path, mimetype, uint(width))
	if err != nil {
		http.Error(writer, "could not get asset version", http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", mimetype)
	writer.Header().Add("Content-Length", fmt.Sprintf("%d", mediaVersion.Size))
	writer.Header().Add("Cache-Control", "public, max-age=604800, immutable")
	_, err = io.Copy(writer, mediaVersion.Content)
	if err != nil {
		log.Println("failed to send media content", err)
		return
	}
}
