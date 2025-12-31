package app

import (
	"strings"
	"testing"
)

func TestColorizeHelpText(t *testing.T) {
	in := "Usage: gifgrep <command> [flags]\n\nFlags:\n  -h, --help  Show help.\n\nCommands:\n  search <query> ... [flags]\n"
	out := colorizeHelpText(in)

	if out == in {
		t.Fatalf("expected colorization")
	}
	if !strings.Contains(out, "\x1b[1mUsage:") {
		t.Fatalf("expected bold Usage heading, got: %q", out)
	}
	if !strings.Contains(out, "\x1b[36m--help\x1b[0m") {
		t.Fatalf("expected colored long flag")
	}
	if !strings.Contains(out, "\x1b[36m-h\x1b[0m") {
		t.Fatalf("expected colored short flag")
	}
	if !strings.Contains(out, "\x1b[36msearch\x1b[0m") {
		t.Fatalf("expected colored command name")
	}
	if !strings.Contains(out, "\x1b[90m<command>\x1b[0m") || !strings.Contains(out, "\x1b[90m[flags]\x1b[0m") {
		t.Fatalf("expected dim placeholders")
	}
}
