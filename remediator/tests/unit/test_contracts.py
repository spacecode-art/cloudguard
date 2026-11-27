"""Contract tests: the same fixtures the Go tests use, validated against the same schemas."""

import json
from pathlib import Path
from typing import Any

import pytest
from jsonschema import Draft202012Validator

# remediator/tests/unit/test_contracts.py -> repository root is three levels up.
SCHEMAS = Path(__file__).resolve().parents[3] / "schemas"
VALID = sorted((SCHEMAS / "fixtures" / "valid").glob("*.json"))
INVALID = sorted((SCHEMAS / "fixtures" / "invalid").glob("*.json"))


def _fixture_id(path: Path) -> str:
    return path.name


def _load(path: Path) -> Any:
    return json.loads(path.read_text(encoding="utf-8"))


def _errors(fixture: Path) -> list[Any]:
    """Validation errors for a fixture; the schema is chosen by filename prefix."""
    prefix = fixture.name.split("-", 1)[0]
    schema = _load(SCHEMAS / f"{prefix}.schema.json")
    Draft202012Validator.check_schema(schema)  # the schema itself must be valid
    return list(Draft202012Validator(schema).iter_errors(_load(fixture)))


def test_fixtures_exist() -> None:
    # Guard against a silent pass: parametrizing over an empty list runs zero tests.
    assert VALID
    assert INVALID


@pytest.mark.parametrize("fixture", VALID, ids=_fixture_id)
def test_valid_fixture_passes(fixture: Path) -> None:
    assert _errors(fixture) == []


@pytest.mark.parametrize("fixture", INVALID, ids=_fixture_id)
def test_invalid_fixture_is_rejected(fixture: Path) -> None:
    assert _errors(fixture) != []
