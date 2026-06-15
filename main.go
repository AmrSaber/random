// Package main is the entry point for the random CLI.
package main

import (
	"github.com/AmrSaber/random/v3/src/cmd"
	"github.com/AmrSaber/random/v3/src/common"
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

	rootOpts := &common.RootOptions{Count: 1}
	rootCmd := cmd.NewRootCommand(v, rootOpts)

	if err := rootCmd.Execute(); err != nil {
		common.Fail(err)
	}
}
