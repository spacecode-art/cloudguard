# ADR-004: JSON Schema for cross-language contracts

- Status: Accepted
- Date: 2026-11-27

## Context

Go produces inventories and drift results. Python consumes them and produces
findings and remediation records. Both sides must agree exactly on structure.

## Problem

How do we define and enforce the data shape shared between Go and Python so
that neither side can silently break the other?

## Decision

**JSON Schema is the authoritative contract.** Schemas live in `schemas/`.
Every document carries a `schema_version`. Both Go and Python validate against
the same schemas and the same shared JSON fixtures in automated tests.
Schema changes are reviewed and versioned, never silently altered.

In the initial release, components exchange these documents as JSON files on
disk. How Go invokes Python is a separate decision, to be recorded in its own
ADR when implemented.

## Alternatives considered

1. **Protocol Buffers / gRPC.** Strong typing and code generation, but adds a
   build step and network service, which contradicts ADR-001.
2. **Types defined in code on one side only.** Fast, but the other language
   drifts out of sync without anyone noticing.
3. **Ad hoc JSON.** Zero setup, no guarantees. Rejected.

## Consequences

- Positive: human-readable, language-neutral, easy to debug and diff.
- Positive: shared fixtures give automated compatibility testing.
- Negative: no compile-time type safety across the boundary; compatibility
  tests in CI are mandatory.
- Negative: schemas must be maintained with discipline.