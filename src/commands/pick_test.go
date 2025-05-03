package commands

import (
	"slices"
	"strings"
	"testing"

	"github.com/AmrSaber/random/v3/src/common"
	"github.com/AmrSaber/random/v3/src/common/tests"

	"github.com/urfave/cli/v2"
)

func TestPickCommand_Basic(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{PickCommand}}
	args := []string{"test-app", "pick", "a", "b", "c"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if !slices.Contains([]string{"a", "b", "c"}, output) {
		t.Errorf("output '%s' is not one of the input items", output)
	}
}

func TestPickCommand_Count(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{PickCommand}}
	args := []string{"test-app", "pick", "--count", "3", "a", "b", "c"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
}

func TestPickCommand_Number(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{PickCommand}}
	args := []string{"test-app", "pick", "--number", "2", "a", "b", "c"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, common.DEFAULT_DELIMITER)
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

func TestPickCommand_Delimiter(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{PickCommand}}
	args := []string{"test-app", "pick", "--delimiter", ",", "--number", "2", "a", "b", "c"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, ",")
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}

	for _, item := range items {
		if !slices.Contains([]string{"a", "b", "c"}, item) {
			t.Errorf("output item '%s' is not one of the input items", item)
		}
	}
}

func TestPickCommand_EmptyInput(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{PickCommand}}
	args := []string{"test-app", "pick"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if output != "" {
		t.Errorf("expected empty output, got '%s'", output)
	}
}
