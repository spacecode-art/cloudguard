"""Exit-code contract. Must stay identical to the Go CLI and the README."""

from typing import Final

EXIT_OK: Final = 0  # successful operation
EXIT_OPERATIONAL_FAILURE: Final = 1  # something went wrong at runtime
EXIT_USAGE: Final = 2  # invalid configuration or usage
EXIT_POLICY_VIOLATIONS: Final = 3  # policy violations detected
EXIT_REMEDIATION_FAILURE: Final = 4  # remediation or verification failure
