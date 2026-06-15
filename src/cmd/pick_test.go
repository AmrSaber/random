package cmd_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/AmrSaber/random/v3/src/common"
	"github.com/AmrSaber/random/v3/src/common/tests"
)

func TestPickCommand_Basic(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"pick", "a", "b", "c"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if !slices.Contains([]string{"a", "b", "c"}, output) {
		t.Fatalf("output %q not expected", output)
	}
}

func TestPickCommand_Count(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"pick", "--count", "3", "a", "b", "c"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}
}

func TestPickCommand_Number(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"pick", "--number", "2", "a", "b", "c"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, common.DefaultDelimiter)
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
}

func TestPickCommand_Delimiter(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"pick", "--delimiter", ",", "--number", "2", "a", "b", "c"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	items := strings.Split(output, ",")
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	for _, item := range items {
		if !slices.Contains([]string{"a", "b", "c"}, item) {
			t.Fatalf("unexpected item %q", item)
		}
	}
}

func TestPickCommand_EmptyInput(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"pick"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	if output := strings.TrimSpace(buf.String()); output != "" {
		t.Fatalf("expected empty output, got %q", output)
	}
}
