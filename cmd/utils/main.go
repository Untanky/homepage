package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/untanky/homepage/internal/cmd/images"
	"github.com/untanky/homepage/internal/cmd/migrate"
	"github.com/untanky/homepage/internal/config"
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
}

func main() {
	runtimeCtx := context.Background()

	ctx := kong.Parse(&CLI,
		kong.Name("utils"),
		kong.Description("Handle utilities for the homepage server"),
		kong.BindTo(runtimeCtx, (*context.Context)(nil)),
	)

	config, err := config.Load(CLI.Config)
	if err != nil {
		fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	ctx.Bind(config)

	if err := ctx.Run(); err != nil {
		fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
