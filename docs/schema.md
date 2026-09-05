# Evidence schema v0

Funding Topology separates **observations** from **inferences**.

An edge is an auditable observation with transaction provenance. The initial PoC exposes:

- `FUNDED_BY`: earliest qualifying native transfer into a seed address.
- `FEE_PAID_BY`: fee payer recorded for a transaction involving a subject address.

Every edge carries a source, target, relationship type, signature, slot, timestamp, and derivation method. Higher-order claims such as common control or operator attribution are intentionally out of scope for the open deterministic core.

Planned grant milestones add `SIGNED_BY`, `CREATED_BY`, `AUTHORIZED_BY`, `DEPLOYED_BY`, transfer/program interaction observations, bounded traversal, and service/high-degree suppression.
