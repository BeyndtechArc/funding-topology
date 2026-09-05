package model

import "time"

type Relationship string

const (
	FundedBy  Relationship = "FUNDED_BY"
	FeePaidBy Relationship = "FEE_PAID_BY"
)

type Provenance struct {
	Signature        string    `json:"signature"`
	Slot             uint64    `json:"slot"`
	Timestamp        time.Time `json:"timestamp"`
	Program          string    `json:"program,omitempty"`
	DerivationMethod string    `json:"derivation_method"`
}

type Edge struct {
	Source       string       `json:"source"`
	Target       string       `json:"target"`
	Relationship Relationship `json:"relationship"`
	Amount       uint64       `json:"amount,omitempty"`
	Provenance   Provenance   `json:"provenance"`
}
