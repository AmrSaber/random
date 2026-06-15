package cmd_test

import (
	"strings"
	"testing"

	"github.com/AmrSaber/random/v3/src/common"
	"github.com/AmrSaber/random/v3/src/common/tests"
)

func TestStringCommand_Basic(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"string"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if len(output) != common.DefaultStringLength {
		t.Fatalf("expected string length %d, got %d", common.DefaultStringLength, len(output))
	}
}

func TestStringCommand_Length(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"string", "--length", "10"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if len(output) != 10 {
		t.Fatalf("expected string length 10, got %d", len(output))
	}
}

func TestStringCommand_Count(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"string", "--count", "3"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d", len(lines))
	}

	for _, line := range lines {
		if len(line) != common.DefaultStringLength {
			t.Fatalf("expected string length %d, got %d", common.DefaultStringLength, len(line))
		}
	}
}

func TestStringCommand_InvalidType(t *testing.T) {
	_, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"string", "--type", "invalid"})

	if err := root.Execute(); err == nil {
		t.Fatal("expected error for invalid type")
	}
}

func TestStringCommand_ValidType(t *testing.T) {
	buf, cleanup := tests.SetupTest()
	defer cleanup()

	opts := &common.RootOptions{}
	root := tests.NewRootCommand(opts)
	root.SetArgs([]string{"string", "--type", "hex"})

	if err := root.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	output := strings.TrimSpace(buf.String())
	if len(output) != common.DefaultStringLength {
		t.Fatalf("expected string length %d, got %d", common.DefaultStringLength, len(output))
	}
}
