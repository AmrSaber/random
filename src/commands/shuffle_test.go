package commands

import (
	"slices"
	"strings"
	"testing"

	"github.com/AmrSaber/random/v3/src/common"
	"github.com/AmrSaber/random/v3/src/common/tests"

	"github.com/urfave/cli/v2"
)

func TestShuffleCommand_Basic(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{ShuffleCommand}}
	input := []string{"a", "b", "c"}
	args := append([]string{"random", "shuffle"}, input...)

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, common.DEFAULT_DELIMITER)
	if len(items) != len(input) {
		t.Errorf("expected %d items, got %d", len(input), len(items))
	}

	// Check all input items are present in output
	for _, item := range input {
		if !slices.Contains(items, item) {
			t.Errorf("item %q not found in output", item)
		}
	}
}

func TestShuffleCommand_Count(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{ShuffleCommand}}
	args := []string{"random", "shuffle", "--count", "3", "a", "b", "c"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}

	// Each line should contain all input items
	for _, line := range lines {
		items := strings.Split(line, common.DEFAULT_DELIMITER)
		if len(items) != 3 {
			t.Errorf("expected 3 items per line, got %d", len(items))
		}
	}
}

func TestShuffleCommand_Delimiter(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{ShuffleCommand}}
	args := []string{"random", "shuffle", "--delimiter", ", ", "a", "b", "c"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, ", ")
	if len(items) != 3 {
		t.Errorf("expected 3 items, got %d", len(items))
	}
}

func TestShuffleCommand_EmptyInput(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{ShuffleCommand}}
	args := []string{"random", "shuffle"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if output != "" {
		t.Errorf("expected empty output, got '%s'", output)
	}
}
