package media

import (
	"fmt"
	"io"
	"mime"
	"slices"
	"sort"
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

var mimetypeOrder = map[string]int{
	"image/avif": -4,
	"image/webp": -3,
	"image/png":  -2,
	"image/jpeg": -1,
}

func Sources(asset Asset, pathPrefix string, width uint) []Source {
	basePath := fmt.Sprintf("%s/%s", pathPrefix, asset.Path)
	sizes := fmt.Sprintf("(min-width: %dpx) %dpx, 100vw", asset.Width, asset.Width)

	srcSets := make(map[string][]string, 3)

	for _, version := range asset.Versions {
		extension, _ := mime.ExtensionsByType(version.Mimetype)
		srcSets[version.Mimetype] = append(srcSets[version.Mimetype], fmt.Sprintf("%s-%d%s %dw", basePath, version.Width, extension[0], version.Width))
	}

	sources := make([]Source, 0, 3)

	for mimetype, sourceSet := range srcSets {
		sources = append(sources, Source{
			Type:      mimetype,
			Sizes:     sizes,
			SourceSet: strings.Join(sourceSet, ", "),
		})
	}

	slices.SortFunc(sources, func(a Source, b Source) int {
		return mimetypeOrder[a.Type] - mimetypeOrder[b.Type]
	})

	return sources
}

func (asset Asset) FallbackURL(pathPrefix string) string {
	if len(asset.Versions) == 0 {
		return ""
	}

	sorted := make([]VersionMetadata, len(asset.Versions))
	copy(sorted, asset.Versions)

	sort.Slice(sorted, func(i, j int) bool {
		pi, pj := mimetypeOrder[sorted[i].Mimetype], mimetypeOrder[sorted[j].Mimetype]
		if pi != pj {
			return pi < pj
		}
		return sorted[i].Width < sorted[i].Height
	})

	extension, _ := mime.ExtensionsByType(sorted[0].Mimetype)
	return fmt.Sprintf("%s/%s-%d%s", pathPrefix, asset.Path, sorted[0].Width, extension[0])
}
