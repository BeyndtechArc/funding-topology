package extract

import "time"

// Transfer is a normalized observation supplied by an RPC/provider adapter.
// The PoC intentionally keeps provider parsing outside the deterministic core.
type Transfer struct {
	From      string
	To        string
	Lamports  uint64
	Signature string
	Slot      uint64
	Timestamp time.Time
}

type Transaction struct {
	Signature string
	Slot      uint64
	Timestamp time.Time
	FeePayer  string
	Signers   []string
}
