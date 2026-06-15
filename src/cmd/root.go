// Package cmd contains the cobra commands for the random CLI.
package cmd

import (
	"github.com/AmrSaber/random/v3/src/common"

	"github.com/spf13/cobra"
)

const countUsage = "Number of results to generate"

// NewRootCommand constructs the root cobra command for the random CLI.
func NewRootCommand(version string, opts *common.RootOptions) *cobra.Command {
	if opts == nil {
		opts = &common.RootOptions{Count: 1}
	}

	if opts.Count == 0 {
		opts.Count = 1
	}

	rootCmd := &cobra.Command{
		Use:           "random",
		Short:         "CLI tool to generate random data",
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       version,
	}

	rootCmd.PersistentFlags().IntVarP(&opts.Count, "count", "c", opts.Count, countUsage)
	rootCmd.AddCommand(
		NewStringCommand(opts),
		NewShuffleCommand(opts),
		NewPickCommand(opts),
		NewIDCommand(opts),
	)

	return rootCmd
}
