// Package cli defines the command-line interface. It parses input and calls
// into other packages; it must not contain cloud, policy or remediation logic.
package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewRootCmd builds the command tree. It is a function (not a global) so tests
// can create a fresh, isolated command every time.
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "cloudguard",
		Short: "Cloud drift detection and compliance remediation for AWS",
		// We print errors ourselves in Execute so the format is consistent.
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          noArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			_ = cmd.Help()
		},
	}

	// Smaller command surface: no auto-generated shell-completion command.
	root.CompletionOptions.DisableDefaultCmd = true

	// Bad flags (e.g. --bogus) are usage errors, exit code 2.
	root.SetFlagErrorFunc(func(_ *cobra.Command, err error) error {
		return &ExitError{Code: ExitUsage, Err: err}
	})

	root.AddCommand(newVersionCmd())

	return root
}

// Execute runs the CLI and returns the process exit code.
func Execute() int {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)

		var exitErr *ExitError
		if errors.As(err, &exitErr) {
			return exitErr.Code
		}
		return ExitOperationalFailure
	}
	return ExitOK
}

// noArgs rejects unexpected positional arguments (including unknown
// subcommands) as usage errors.
func noArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.NoArgs(cmd, args); err != nil {
		return &ExitError{Code: ExitUsage, Err: err}
	}
	return nil
}
