package commands

import (
	"strings"
	"testing"

	"github.com/AmrSaber/random/src/common"
	"github.com/AmrSaber/random/src/common/tests"

	"github.com/urfave/cli/v2"
)

func TestStringCommand_Basic(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{StringCommand}}
	args := []string{"random", "string"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if len(output) != common.DEFAULT_STRING_LENGTH {
		t.Errorf("expected string length %d, got %d", common.DEFAULT_STRING_LENGTH, len(output))
	}
}

func TestStringCommand_Length(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{StringCommand}}
	args := []string{"random", "string", "--length", "10"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if len(output) != 10 {
		t.Errorf("expected string length 10, got %d", len(output))
	}
}

func TestStringCommand_Count(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{StringCommand}}
	args := []string{"random", "string", "--count", "3"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}

	for _, line := range lines {
		if len(line) != common.DEFAULT_STRING_LENGTH {
			t.Errorf("expected string length %d, got %d", common.DEFAULT_STRING_LENGTH, len(line))
		}
	}
}

func TestStringCommand_InvalidType(t *testing.T) {
	_, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{StringCommand}}
	args := []string{"random", "string", "--type", "invalid"}

	if err := app.Run(args); err == nil {
		t.Error("expected error for invalid type, got nil")
	}
}

func TestStringCommand_ValidType(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	app := &cli.App{Commands: []*cli.Command{StringCommand}}
	args := []string{"random", "string", "--type", "hex"}

	if err := app.Run(args); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if len(output) != common.DEFAULT_STRING_LENGTH {
		t.Errorf("expected string length %d, got %d", common.DEFAULT_STRING_LENGTH, len(output))
	}
}
