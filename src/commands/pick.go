package commands

import (
	"strings"

	"github.com/AmrSaber/random/src/common"
	"github.com/AmrSaber/random/src/common/helpers"

	"github.com/urfave/cli/v2"
)

var PickCommand = &cli.Command{
	Name:        "pick",
	Description: "Picks and print a random item from the given items",

	Flags: []cli.Flag{
		common.CountFlag,

		&cli.StringFlag{
			Name:    "delimiter",
			Aliases: []string{"d"},
			Usage:   "Delimiter for the shuffled output",
			Value:   common.DEFAULT_DELIMITER,
		},

		&cli.IntFlag{
			Name:    "number",
			Aliases: []string{"n"},
			Usage:   "Number of items to pick",
			Value:   1,
		},
	},

	Action: func(ctx *cli.Context) error {
		count := ctx.Int("count")
		delimiter := ctx.String("delimiter")
		number := ctx.Int("number")
		values := ctx.Args().Slice()

		if len(values) == 0 {
			return nil
		}

		randomChoices := make([]string, 0, count)

		for range count {
			shuffledItems := helpers.Shuffle(values)
			chosenItems := shuffledItems[:number]
			randomChoices = append(randomChoices, strings.Join(chosenItems, delimiter))
		}

		common.Std.Println(strings.Join(randomChoices, "\n"))

		return nil
	},
}
