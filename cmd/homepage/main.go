package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"go.lukasgrimm.me/homepage/internal/cmd/images"
	"go.lukasgrimm.me/homepage/internal/cmd/migrate"
	"go.lukasgrimm.me/homepage/internal/cmd/posts"
	"go.lukasgrimm.me/homepage/internal/cmd/server"
	"go.lukasgrimm.me/homepage/internal/config"
)

var CLI struct {
	Config string `help:"Path to the config path" type:"existingFile"`
	Images struct {
		Get    images.GetCommand    `cmd:"" help:"Get an images"`
		List   images.ListCommand   `cmd:"" help:"List images"`
		Create images.CreateCommand `cmd:"" help:"Create an image"`
	} `cmd:"" group:"Images" help:"Manage images"`
	Migrations struct {
		Up migrate.Up `cmd:"" help:"Run up migrations"`
	} `cmd:"" group:"Migrations" help:"Manage database migrations"`
	Posts posts.Group         `cmd:"" group:"Posts" help:"Manage blog posts"`
	Serve server.ServeCommand `cmd:"" help:"Start the homepage server"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("utils"),
		kong.Description("Handle utilities for the homepage server"),
	)

	config, err := config.Load(CLI.Config)
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	runtimeCtx := context.Background()

	ctx.BindTo(runtimeCtx, (*context.Context)(nil))
	ctx.Bind(config)

	if err := ctx.Run(); err != nil {
		fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
