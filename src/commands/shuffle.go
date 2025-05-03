package commands

import (
	"strings"

	"github.com/AmrSaber/random/src/common"
	"github.com/AmrSaber/random/src/common/helpers"

	"github.com/urfave/cli/v2"
)

var ShuffleCommand = &cli.Command{
	Name:        "shuffle",
	Description: "Shuffles the given items",

	Flags: []cli.Flag{
		common.CountFlag,

		&cli.StringFlag{
			Name:    "delimiter",
			Aliases: []string{"d"},
			Usage:   "Delimiter for the shuffled output",
			Value:   common.DEFAULT_DELIMITER,
		},
	},

	Action: func(ctx *cli.Context) error {
		count := ctx.Int("count")
		delimiter := ctx.String("delimiter")
		values := ctx.Args().Slice()

		if len(values) == 0 {
			return nil
		}

		randomShuffles := make([]string, 0, count)

		for range count {
			shuffledItems := helpers.Shuffle(values)
			randomShuffles = append(randomShuffles, strings.Join(shuffledItems, delimiter))
		}

		// Print the generated random strings
		common.Std.Println(strings.Join(randomShuffles, "\n"))

		return nil
	},
}
