# Validation cases

Funding Topology is being developed against real Solana investigations rather than synthetic-only examples. The cases below define the initial reproducibility targets for the open-source primitive.

## KYLIE — existing methodology demonstration

Rugburn's published KYLIE investigation showed why pre-event operational context matters. The investigation reconstructed funding and early-buyer infrastructure that existed before a compromised X account promoted a Solana token.

Funding Topology does **not** claim that the current PoC can reproduce the full published investigation yet. Instead, KYLIE is the first benchmark for converting already-proven manual investigative steps into deterministic, provenance-backed relationships.

Initial reproducibility goals:

- recover the earliest qualifying funding observations for selected case wallets;
- preserve transaction signature, slot, timestamp, amount, and derivation method for each edge;
- recover fee-payer relationships where available;
- compare machine-produced evidence edges with the relationships documented during the original investigation;
- record any relationship that still requires analyst inference rather than representing it as deterministic fact.

## Aquifer — next reproducible target

Aquifer is the next adversarial validation target. It is useful because the attacker infrastructure includes a fresh/disposable execution wallet and a short preparation window, which makes conventional address-history heuristics weaker and tests whether operational context can still be recovered from deterministic evidence.

The initial open-source reproduction will begin from public seed addresses and attempt to reconstruct:

1. the first qualifying funding relationship into the attacker wallet;
2. the transaction fee payer associated with relevant preparation activity;
3. signer and account-creation observations as those primitives land;
4. deployment/authority relationships for malicious infrastructure where deterministically recoverable;
5. bounded one-hop pivots with explicit service/high-degree-node suppression.

The objective is **not** to publish an operator identity. The objective is to determine which operational relationships can be independently reproduced from public Solana history, which require additional data, and which remain inference.

## Follow-on validation

After KYLIE and Aquifer, a third case will be selected to stress fragmented or multi-stage infrastructure. Rain is a candidate because its public incident history involves multiple operational steps and cross-chain exit behavior, but it is intentionally outside the first PoC so the initial grant scope stays narrow.

## Evaluation criteria

For every case, validation records:

- deterministic edges recovered;
- transaction-level provenance completeness;
- false-positive relationships;
- service/high-degree-node suppression behavior;
- missing-data/provider dependencies;
- manual analyst steps still required;
- conclusions that remain inference rather than fact.
