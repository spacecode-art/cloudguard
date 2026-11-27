// Package version holds build metadata. The values are injected at build time
// with -ldflags, so a released binary can always report exactly what it is.
package version

// These are variables (not constants) because the linker overwrites them.
var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)
