# ADR-001: Modular monolith architecture

- Status: Accepted
- Date: 2026-11-27

## Context

CloudGuard has several distinct responsibilities: discovery, drift detection,
policy evaluation, remediation planning, execution, verification and audit.
It is built by a small team (initially one engineer) with an approximately
eight-week target for the first release.

## Problem

How should the system be divided into deployable units?

## Decision

CloudGuard is a **modular monolith**: one repository and one deployable unit
(a Go binary plus a Python package, shipped in a single container image).
Components have strict internal boundaries and communicate only through
versioned contracts (see ADR-004).

Privilege separation is **not** provided by process boundaries. It is enforced
by distinct AWS IAM identities (Scanner, Remediator, Deployment), because
application-level checks alone are not a security boundary.

## Alternatives considered

1. **Microservices.** Independent scaling and deployment, but adds networking,
   service authentication, distributed tracing and failure modes that the
   current scale does not justify.
2. **Monolith without enforced boundaries.** Simplest to start, but discovery
   and mutation logic would blur together, which is a security risk.
3. **Serverless functions per stage.** Cheap when idle, but harder to test
   locally and ties the design to one cloud provider's runtime.

## Consequences

- Positive: simple build, test and deploy; easy local development; low cost.
- Positive: boundaries can later be split into services if needed.
- Negative: components cannot scale independently.
- Negative: boundaries rely on engineering discipline (reviews, tests) rather
  than the network.
- Revisit if CloudGuard becomes multi-tenant or needs independent scaling.