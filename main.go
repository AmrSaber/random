package main

import (
	"os"

	"github.com/AmrSaber/random/v3/src/commands"
	"github.com/AmrSaber/random/v3/src/common"

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
