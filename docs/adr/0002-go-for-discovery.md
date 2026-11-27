# ADR-002: Go for resource discovery

- Status: Accepted
- Date: 2026-11-27

## Context

Discovery makes many concurrent, read-only calls to AWS APIs. It must handle
pagination, throttling, retries and partial failures, and it ships as a CLI.

## Problem

Which language should implement the CLI, collectors, normalization and drift
detection?

## Decision

Use **Go** with the AWS SDK for Go v2 and Cobra for the CLI.

Reasons:
- Goroutines, channels and `context` make bounded concurrency and cancellation
  first-class.
- Compiles to a single static binary, which suits a minimal, non-root
  container image.
- Statically typed, with a strong standard library and testing tooling.
- The AWS SDK for Go v2 is officially maintained.

## Alternatives considered

1. **Python for everything.** One language, but weaker concurrency ergonomics
   and a heavier runtime to distribute.
2. **Rust.** Excellent safety and performance, but a steeper learning curve
   and slower iteration for this project's goals.
3. **Java or Kotlin.** Mature AWS tooling, but heavier runtime and
   build footprint for a CLI.

## Consequences

- Positive: predictable performance, simple distribution, safe concurrency.
- Negative: more verbose than Python; explicit error handling everywhere.
- Negative: two language toolchains to maintain (see ADR-003).