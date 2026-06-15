package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AmrSaber/random/v3/src/common"

	"github.com/spf13/cobra"
)

// NewStringCommand wires the string subcommand.
func NewStringCommand(opts *common.RootOptions) *cobra.Command {
	var length int
	var strType string

	cmd := &cobra.Command{
		Use:     "string",
		Aliases: []string{"str"},
		Short:   "Prints a random string",
		RunE: func(_ *cobra.Command, _ []string) error {
			if !slices.Contains(common.StringTypes, strType) {
				return fmt.Errorf("invalid string type %q", strType)
			}

			count := opts.Count
			randomStrings := make([]string, 0, count)

			for range count {
				randomStrings = append(randomStrings, common.GetRandomString(strType, length))
			}

			common.Std.Println(strings.Join(randomStrings, "\n"))
			return nil
		},
	}

	cmd.Flags().IntVarP(&length, "length", "l", common.DefaultStringLength, "Length of the random string")
	cmd.Flags().StringVarP(&strType, "type", "t", common.StringTypeASCII, fmt.Sprintf("Type of the random string, must be one of [%s]", strings.Join(common.StringTypes, ", ")))

	return cmd
}
