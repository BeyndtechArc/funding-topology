package extract

import (
	"errors"

	"github.com/BeyndtechArc/funding-topology/internal/model"
)

// FeePayerEdge records a transaction-level fee-payer relationship.
// It deliberately does not infer common control between signer and fee payer.
func FeePayerEdge(subject string, tx Transaction) (model.Edge, error) {
	if subject == "" || tx.FeePayer == "" || tx.Signature == "" {
		return model.Edge{}, errors.New("missing required transaction evidence")
	}
	return model.Edge{
		Source:       subject,
		Target:       tx.FeePayer,
		Relationship: model.FeePaidBy,
		Provenance: model.Provenance{
			Signature:        tx.Signature,
			Slot:             tx.Slot,
			Timestamp:        tx.Timestamp,
			DerivationMethod: "transaction_fee_payer_v0",
		},
	}, nil
}
