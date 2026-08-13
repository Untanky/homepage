package main

import (
	"context"
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
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

	if err := ctx.Run(); err != nil {
		fmt.Fprintf(ctx.Stderr, "Error: %v\n", err)
	}
}
