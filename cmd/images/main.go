package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/untanky/homepage/internal/config"
)

var CLI struct {
	Config string        `help:"Path to the config path" type:"existingFile"`
	Get    getCommand    `cmd:"" help:"Get an images"`
	List   listCommand   `cmd:"" help:"List images"`
	Create createCommand `cmd:"" help:"Create an image"`
}

func main() {
	runtimeCtx := context.Background()

	ctx := kong.Parse(&CLI,
		kong.Name("images"),
		kong.Description("Manage images"),
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
