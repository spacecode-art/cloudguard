package cli

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestVersionCommandPrintsProgramName(t *testing.T) {
	var out bytes.Buffer
	root := NewRootCmd()
	root.SetOut(&out)
	root.SetArgs([]string{"version"})

	if err := root.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(out.String(), "cloudguard ") {
		t.Errorf("unexpected output: %q", out.String())
	}
}

func TestUsageErrorsReturnExitUsage(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"unknown command", []string{"nonsense"}},
		{"unknown flag", []string{"version", "--bogus"}},
		{"unexpected argument", []string{"version", "extra"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := NewRootCmd()
			root.SetOut(io.Discard)
			root.SetErr(io.Discard)
			root.SetArgs(tc.args)

			err := root.Execute()

			var exitErr *ExitError
			if !errors.As(err, &exitErr) {
				t.Fatalf("expected *ExitError, got %v", err)
			}
			if exitErr.Code != ExitUsage {
				t.Errorf("exit code = %d, want %d", exitErr.Code, ExitUsage)
			}
		})
	}
}
