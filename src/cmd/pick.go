package cmd

import (
	"strings"

	"github.com/AmrSaber/random/v3/src/common"

	"github.com/spf13/cobra"
)

// NewPickCommand wires the pick subcommand.
func NewPickCommand(opts *common.RootOptions) *cobra.Command {
	var delimiter string
	var number int

	cmd := &cobra.Command{
		Use:   "pick [items...]",
		Short: "Picks and print a random item from the given items",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			count := opts.Count
			randomChoices := make([]string, 0, count)

			for range count {
				shuffledItems := common.Shuffle(args)
				chosenItems := shuffledItems[:number]
				randomChoices = append(randomChoices, strings.Join(chosenItems, delimiter))
			}

			common.Std.Println(strings.Join(randomChoices, "\n"))
			return nil
		},
	}

	cmd.Flags().StringVarP(&delimiter, "delimiter", "d", common.DefaultDelimiter, "Delimiter for the shuffled output")
	cmd.Flags().IntVarP(&number, "number", "n", 1, "Number of items to pick")

	return cmd
}
