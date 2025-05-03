package main

import (
	"os"
	"random/src/commands"
	"random/src/common"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "random",
		Description: "CLI tool to generate random data",
		Commands:    []*cli.Command{commands.StringCommand, commands.ShuffleCommand, commands.PickCommand, commands.IdCommand},
	}

	// Run CLI
	if err := app.Run(os.Args); err != nil {
		common.Fail(err)
	}
}
