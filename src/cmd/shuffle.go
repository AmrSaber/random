package cmd

import (
	"strings"

	"github.com/AmrSaber/random/v3/src/common"

	"github.com/spf13/cobra"
)

// NewShuffleCommand wires the shuffle subcommand.
func NewShuffleCommand(opts *common.RootOptions) *cobra.Command {
	var delimiter string

	cmd := &cobra.Command{
		Use:   "shuffle [items...]",
		Short: "Shuffles the given items",
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				return nil
			}

			count := opts.Count
			randomShuffles := make([]string, 0, count)

			for range count {
				shuffledItems := common.Shuffle(args)
				randomShuffles = append(randomShuffles, strings.Join(shuffledItems, delimiter))
			}

			common.Std.Println(strings.Join(randomShuffles, "\n"))
			return nil
		},
	}

	cmd.Flags().StringVarP(&delimiter, "delimiter", "d", common.DefaultDelimiter, "Delimiter for the shuffled output")

	return cmd
}
