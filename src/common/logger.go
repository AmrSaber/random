package common

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// Err writes colored error output for CLI feedback.
var Err = color.New(color.FgRed)

// Std captures standard output for CLI commands.
var Std = log.New(os.Stdout, "", 0)

// Fail prints the error message and terminates the process.
func Fail(message any) {
	_, _ = Err.Fprint(os.Stderr, fmt.Sprintf("%s\n", message))
	os.Exit(1)
}
