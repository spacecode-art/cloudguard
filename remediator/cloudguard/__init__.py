"""CloudGuard Python component: policy evaluation, remediation, verification and audit."""

from importlib.metadata import PackageNotFoundError, version

try:
    # Single source of truth for the version is pyproject.toml.
    __version__ = version("cloudguard-remediator")
except PackageNotFoundError:  # running from a source tree that was never installed
    __version__ = "0.0.0+unknown"
