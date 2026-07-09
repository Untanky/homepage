package sql

import (
	"bytes"
	"io"

	"github.com/untanky/homepage/blog"
)

type memoryPost struct {
	metadata blog.PostMetadata
	content  *bytes.Buffer
}

func (post *memoryPost) Metadata() blog.PostMetadata {
	return post.metadata
}

func (post *memoryPost) WriteTo(writer io.Writer) (int64, error) {
	return post.content.WriteTo(writer)
}

func (post *memoryPost) Close() error {
	return nil
}
