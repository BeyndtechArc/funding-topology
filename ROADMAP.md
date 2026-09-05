# Roadmap

This repository is an early public proof-of-concept for a proposed Solana developer-tooling grant.

## M1 — Evidence schema + deterministic extractors
- Versioned provenance schema
- `FUNDED_BY`
- `FEE_PAID_BY`
- `SIGNED_BY`
- account/deployment/authority observations
- fixture-driven tests

## M2 — Bounded topology traversal + suppression
- depth/time/edge constraints
- cycle detection
- traversal budgets
- deterministic ordering
- high-degree/service-origin suppression hooks

## M3 — Incident validation corpus
- KYLIE: reproduce deterministic portions of a previously completed Rugburn investigation
- Aquifer: next fresh/disposable-wallet reproducibility target
- one additional public Solana incident class after the first two benchmarks are stable
- false-positive analysis
- missing-data behavior
- explicit inference boundaries

## M4 — Public developer release
- installable CLI/reference library
- JSON/JSONL output
- docs/examples/CI
- six months maintenance and measurable adoption milestones
