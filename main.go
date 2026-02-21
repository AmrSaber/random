package main

import (
	"os"

	"github.com/AmrSaber/random/v3/src/commands"
	"github.com/AmrSaber/random/v3/src/common"

	"github.com/urfave/cli/v2"
)

var version string

func main() {
	if version != "" {
		common.SetVersion(version)
	}

	v := common.GetVersion()
	if v == "" {
		v = "??"
	}

	app := &cli.App{
		Name:        "random",
		Description: "CLI tool to generate random data",
		Version:     v,
		Commands:    []*cli.Command{commands.StringCommand, commands.ShuffleCommand, commands.PickCommand, commands.IdCommand},
	}

	// Run CLI
	if err := app.Run(os.Args); err != nil {
		common.Fail(err)
	}
}
