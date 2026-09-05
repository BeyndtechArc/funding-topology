# Funding Topology

**Provenance-backed operational relationship infrastructure for Solana investigations.**

Funding Topology is an early open-source developer-tooling project from [Rugburn](https://rugburn.io). It is designed to turn raw Solana transaction observations into typed, reproducible evidence relationships while keeping deterministic facts separate from operator-attribution hypotheses.

> Status: early public PoC. This repository does **not** represent a finished production system.

## Why this exists

Solana tooling already makes it possible to retrieve transaction history, decode instructions, inspect labels, and trace transfers. Investigators still repeatedly rebuild a lower-level evidence layer by hand: initial funding, fee-payer relationships, signer/deployer/authority activity, account preparation, temporal sequencing, and the exact transaction provenance behind every claim.

Funding Topology aims to make that evidence stage reusable.

## Built from live investigations

This project is being extracted from Rugburn investigations, not designed from synthetic graph examples alone.

- **KYLIE — methodology already demonstrated manually.** Rugburn reconstructed funding and early-buyer infrastructure that existed before a compromised X account promoted a Solana token. Funding Topology uses that completed investigation as the first benchmark for converting manual evidence work into reproducible primitives.
- **Aquifer — next reproducible target.** Aquifer is the next adversarial test because the attacker infrastructure includes a fresh/disposable wallet and a short preparation window. The target is to reproduce first funding, fee-payer, signer/account-creation and deployment relationships from public Solana history without turning correlation into attribution.
- **Rain — follow-on candidate.** Reserved for a later validation pass once the deterministic core and suppression rules are stable.

See [`docs/cases.md`](docs/cases.md) for the validation contract and boundaries.

## Current PoC

The first commit intentionally implements only two deterministic primitives:

- `FUNDED_BY` — earliest qualifying native funding transfer for a seed address.
- `FEE_PAID_BY` — transaction-level fee-payer relationship.

Both return typed edges with signature/slot/timestamp provenance. Neither primitive asserts common ownership or operator identity.

```text
raw normalized observations
        │
        ▼
deterministic evidence edges
        │
        ▼
bounded topology traversal   (planned)
        │
        ▼
downstream inference / attribution   (out of scope here)
```

## Design principles

1. **Evidence before inference.** Every deterministic edge must be reproducible from chain data.
2. **Provenance is mandatory.** Relationships carry the transaction evidence and derivation method that produced them.
3. **No garbage graph.** Exchange/bridge/relay/high-degree service nodes need explicit handling before topology expansion is meaningful.
4. **Provider portability.** RPC/indexing vendors belong behind adapters; the evidence model should not depend on one vendor.
5. **False positives are a first-class output.** Validation must document what cannot be safely inferred.

## Layout

- `internal/model` — relationship and provenance types
- `internal/extract` — deterministic observation extractors
- `docs/schema.md` — evidence semantics
- `docs/cases.md` — real-incident validation targets (KYLIE → Aquifer → follow-on)
- `ROADMAP.md` — grant-aligned milestones
- `cmd/funding-topology` — CLI seed

## Run tests

```bash
go test ./...
```

## Relationship to Rugburn

The grant-funded Funding Topology project is intended to be completely public and open-source. Rugburn's separate commercial products and higher-order intelligence—cross-incident behavioral clustering, operator attribution, Atlas, Augury, and commercial API/MCP infrastructure—are outside this repository's scope.

## License

Apache-2.0.
