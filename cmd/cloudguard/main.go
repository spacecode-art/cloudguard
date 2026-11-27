// Command cloudguard is the CloudGuard CLI entry point.
package main

import (
	"os"

	"github.com/spacecode-art/cloudguard/internal/cli"
)

func main() {
	// os.Exit skips deferred calls, so main does nothing else.
	os.Exit(cli.Execute())
}
