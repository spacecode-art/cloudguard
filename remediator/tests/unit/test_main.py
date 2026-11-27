import pytest

from cloudguard import __version__
from cloudguard.exitcodes import EXIT_OK, EXIT_USAGE
from cloudguard.main import main


def test_version_flag_prints_version_and_exits_ok(capsys: pytest.CaptureFixture[str]) -> None:
    code = main(["--version"])

    assert code == EXIT_OK
    assert __version__ in capsys.readouterr().out


def test_no_arguments_prints_help_and_exits_ok(capsys: pytest.CaptureFixture[str]) -> None:
    code = main([])

    assert code == EXIT_OK
    assert "usage:" in capsys.readouterr().out


@pytest.mark.parametrize("args", [["--bogus"], ["unexpected"]])
def test_invalid_usage_exits_with_usage_code(args: list[str]) -> None:

    assert main(args) == EXIT_USAGE
