package cmd_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/AmrSaber/random/v3/src/common"
	"github.com/AmrSaber/random/v3/src/common/tests"
)

func TestShuffleCommand_Basic(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	input := []string{"a", "b", "c"}
	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs(append([]string{"shuffle"}, input...))

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, common.DefaultDelimiter)
	if len(items) != len(input) {
		t.Fatalf("expected %d items, got %d", len(input), len(items))
	}

	for _, item := range input {
		if !slices.Contains(items, item) {
			t.Fatalf("missing item %q", item)
		}
	}
}

func TestShuffleCommand_Count(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"shuffle", "--count", "3", "a", "b", "c"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}

	for _, line := range lines {
		items := strings.Split(line, common.DefaultDelimiter)
		if len(items) != 3 {
			t.Fatalf("expected 3 items, got %d", len(items))
		}
	}
}

func TestShuffleCommand_Delimiter(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"shuffle", "--delimiter", ", ", "a", "b", "c"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, ", ")
	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}
}

func TestShuffleCommand_EmptyInput(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"shuffle"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	if output := strings.TrimSpace(buf.String()); output != "" {
		t.Fatalf("expected empty output, got %q", output)
	}
}
