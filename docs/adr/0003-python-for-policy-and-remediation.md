# ADR-003: Python for policy evaluation and remediation

- Status: Accepted
- Date: 2026-11-27

## Context

Policy evaluation and remediation need fast iteration, strong data validation,
and mature AWS mutation tooling. Policies are declarative YAML written and
reviewed by humans.

## Problem

Which language should implement the policy engine, risk classifier,
remediation planner and executor, verification and audit?

## Decision

Use **Python** with Pydantic (validation), boto3 (AWS), pytest, Ruff and mypy.

Policies are YAML and are evaluated by a **fixed set of operators** (`equals`,
`not_equals`, `contains`, `greater_than`, `exists`). Policy files are parsed
with a safe YAML loader and are never executed as code.

## Alternatives considered

1. **Go for everything.** One toolchain, but slower to iterate on rule logic
   and less convenient for validation-heavy models.
2. **OPA/Rego.** Purpose-built policy language, but adds a new language and
   runtime. Kept as an optional future integration.
3. **Python with `eval` or embedded expressions in policies.** Flexible, but
   rejected outright: arbitrary code execution from policy files is unsafe.

## Consequences

- Positive: fast development, expressive validation, mature boto3 support.
- Positive: deterministic, auditable policy evaluation.
- Negative: dynamic typing, mitigated by type hints, mypy and Pydantic.
- Negative: two toolchains and a language boundary, managed through ADR-004.