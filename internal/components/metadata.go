package components

import (
	"fmt"
	"time"

	"github.com/untanky/homepage/blog"
)

type PageMetadata struct {
	Title        string
	Description  string
	CanonicalURL string
	Robots       string
	Author       string
	Type         string
	Image        string
	Site         string
	PublishedAt  time.Time
	EditedAt     time.Time
}

func NewMetadataFromPostMetadata(metadata blog.PostMetadata, blg blog.Blog) PageMetadata {
	return PageMetadata{
		Title:        metadata.Title,
		Description:  metadata.Summary,
		CanonicalURL: fmt.Sprintf("%s/%s", blg.Authority, metadata.Slug),
		Robots:       "index, follow",
		Author:       metadata.Author.Name,
		Type:         "article",
		Image:        fmt.Sprintf("%s%s", blg.Authority, metadata.Banner.FallbackURL(pathPrefix)),
		Site:         blg.Title,
		PublishedAt:  metadata.CreatedAt,
		EditedAt:     metadata.UpdatedAt,
	}
}
