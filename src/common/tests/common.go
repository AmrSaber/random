// Package tests contains helpers for CLI command testing.
package tests

import (
	"bytes"
	"log"

	"github.com/AmrSaber/random/v3/src/cmd"
	"github.com/AmrSaber/random/v3/src/common"

	"github.com/spf13/cobra"
)

// SetupTest redirects command output to a buffer and returns a cleanup function.
func SetupTest() (*bytes.Buffer, func()) {
	var buf bytes.Buffer
	oldStd := common.Std
	common.Std = log.New(&buf, "", 0)
	return &buf, func() { common.Std = oldStd }
}

// NewRootCommand constructs a root command preconfigured for tests.
func NewRootCommand(opts *common.RootOptions) *cobra.Command {
	if opts.Count == 0 {
		opts.Count = 1
	}

	root := cmd.NewRootCommand("", opts)
	root.SilenceErrors = true
	root.SilenceUsage = true

	return root
}
