package media

import (
	"fmt"
	"io"
	"mime"
	"strings"

	"github.com/google/uuid"
)

type (
	AssetID uuid.UUID
)

type Asset struct {
	ID       AssetID
	Path     string
	Alt      string
	Width    uint
	Height   uint
	Versions []VersionMetadata
}

type VersionMetadata struct {
	Mimetype string
	Width    uint
	Height   uint
}

type Version struct {
	AssetID AssetID
	VersionMetadata
	Size    int
	Content io.ReadSeeker
}

type Source struct {
	Type      string
	Sizes     string
	SourceSet string
}

func Sources(asset Asset, pathPrefix string, width uint) []Source {
	basePath := fmt.Sprintf("/%s/%s", pathPrefix, asset.Path)
	sizes := fmt.Sprintf("(min-width: %dpx) %dpx, 100vw")

	srcSets := make(map[string][]string, 3)

	for _, version := range asset.Versions {
		extension, _ := mime.ExtensionsByType(version.Mimetype)
		srcSets[version.Mimetype] = append(srcSets[version.Mimetype], fmt.Sprintf("%s-@%d%s %dw", basePath, version.Width, extension[0], version.Width))
	}

	sources := make([]Source, 0, 3)

	for mimetype, sourceSet := range srcSets {
		sources = append(sources, Source{
			Type:      mimetype,
			Sizes:     sizes,
			SourceSet: strings.Join(sourceSet, ", "),
		})
	}

	return sources
}
