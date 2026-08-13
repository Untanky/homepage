package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Get    getCommand    `cmd:"" help:"Get an images"`
	List   listCommand   `cmd:"" help:"List images"`
	Create createCommand `cmd:"" help:"Create an image"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("images"),
		kong.Description("Manage images"),
	)

	if err := ctx.Run(ctx); err != nil {
		fmt.Fprintf(ctx.Stderr, "Error: %v", err)
	}
}
