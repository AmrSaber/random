package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AmrSaber/random/v3/src/common"
	"github.com/AmrSaber/random/v3/src/common/helpers"

	"github.com/urfave/cli/v2"
)

var StringCommand = &cli.Command{
	Name:        "string",
	Aliases:     []string{"str"},
	Description: "Prints a random string",

	Flags: []cli.Flag{
		common.CountFlag,

		&cli.IntFlag{
			Name:    "length",
			Aliases: []string{"l"},
			Usage:   "Length of the random string",
			Value:   common.DEFAULT_STRING_LENGTH,
		},

		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   fmt.Sprintf("Type of the random string, must be one of [%s]", strings.Join(common.STRING_TYPES, ", ")),
			Value:   common.STRING_TYPE_ASCII,
			Action: func(ctx *cli.Context, s string) error {
				if !slices.Contains(common.STRING_TYPES, s) {
					return fmt.Errorf("invalid string type %q", s)
				}

				return nil
			},
		},
	},

	Action: func(ctx *cli.Context) error {
		count := ctx.Int("count")
		length := ctx.Int("length")
		strType := ctx.String("type")

		randomStrings := make([]string, 0, count)

		for range count {
			randomStrings = append(randomStrings, helpers.GetRandomString(strType, length))
		}

		// Print the generated random strings
		common.Std.Println(strings.Join(randomStrings, "\n"))

		return nil
	},
}
