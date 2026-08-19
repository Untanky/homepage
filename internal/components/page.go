package components

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type HtmlPage struct {
	title    string
	metadata PageMetadata

	scripts     []string
	stylesheets []string

	child templ.Component
}

type PageOption func(cfg *HtmlPage)

func WithTitle(title string) PageOption {
	return func(cfg *HtmlPage) {
		cfg.title = title
	}
}

func WithMetadata(md PageMetadata) PageOption {
	return func(page *HtmlPage) {
		page.metadata = md
	}
}

func WithScript(scriptURL string) PageOption {
	return func(cfg *HtmlPage) {
		cfg.scripts = append(cfg.scripts, scriptURL)
	}
}

func WithStylesheet(stylesheetURL string) PageOption {
	return func(cfg *HtmlPage) {
		cfg.stylesheets = append(cfg.stylesheets, stylesheetURL)
	}
}

func WithChild(child templ.Component) PageOption {
	return func(cfg *HtmlPage) {
		cfg.child = child
	}
}

func Page(options ...PageOption) HtmlPage {
	cfg := HtmlPage{}
	for _, opt := range options {
		opt(&cfg)
	}

	return cfg
}

func (page HtmlPage) Render(ctx context.Context, writer io.Writer) error {
	ctx = templ.WithChildren(ctx, page.child)

	return html(page).Render(ctx, writer)
}
