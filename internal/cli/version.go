package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/spacecode-art/cloudguard/internal/version"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display version and build information",
		Args:  noArgs,
		Run: func(cmd *cobra.Command, _ []string) {
			// Write to cmd's output (not fmt.Println) so tests can capture it.
			fmt.Fprintf(cmd.OutOrStdout(),
				"cloudguard %s (commit %s, built %s, %s %s/%s)\n",
				version.Version, version.Commit, version.Date,
				runtime.Version(), runtime.GOOS, runtime.GOARCH,
			)
		},
	}
}
