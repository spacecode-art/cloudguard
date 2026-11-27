"""Entry point for the CloudGuard Python component."""

import argparse
import sys
from collections.abc import Sequence

from cloudguard import __version__
from cloudguard.exitcodes import EXIT_OK, EXIT_USAGE


def build_parser() -> argparse.ArgumentParser:
    """Create the argument parser (a function, so tests get a fresh one each time)."""
    parser = argparse.ArgumentParser(
        prog="cloudguard-remediator",
        description="CloudGuard policy evaluation and remediation (Python component).",
    )
    parser.add_argument(
        "--version",
        action="version",
        version=f"cloudguard-remediator {__version__}",
    )
    return parser


def main(argv: Sequence[str] | None = None) -> int:
    """Run the CLI and return the process exit code. Never calls sys.exit itself."""
    parser = build_parser()
    try:
        parser.parse_args(argv)
    except SystemExit as exc:
        # argparse exits by raising SystemExit: 0 for --version/--help, 2 for bad usage.
        return exc.code if isinstance(exc.code, int) else EXIT_USAGE

    parser.print_help()
    return EXIT_OK


if __name__ == "__main__":
    sys.exit(main())
