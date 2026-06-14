package components

import (
	"context"
	"fmt"
	"io"

	"github.com/untanky/homepage/blog"
)

type writerToComponent struct {
	writer io.WriterTo
}

func post(post blog.Post) *writerToComponent {
	return new(writerToComponent{
		writer: post,
	})
}

func (component *writerToComponent) Render(ctx context.Context, writer io.Writer) error {
	if _, err := component.writer.WriteTo(writer); err != nil {
		return fmt.Errorf("streaming content: %w", err)
	}

	return nil
}
