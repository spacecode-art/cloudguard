package cli

// Exit-code contract (documented in README.md). Scripts and CI rely on these
// values, so changing them is a breaking change.
const (
	ExitOK                 = 0 // successful operation
	ExitOperationalFailure = 1 // something went wrong at runtime
	ExitUsage              = 2 // invalid configuration or usage
	ExitPolicyViolations   = 3 // policy violations detected
	ExitRemediationFailure = 4 // remediation or verification failure
)

// ExitError carries a specific exit code up to main. Commands return it when
// the default code (1) would be wrong.
type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string { return e.Err.Error() }

// Unwrap lets errors.Is / errors.As see the underlying error.
func (e *ExitError) Unwrap() error { return e.Err }
