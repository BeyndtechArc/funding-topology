package extract

import (
	"errors"
	"sort"

	"github.com/BeyndtechArc/funding-topology/internal/model"
)

// EarliestFunding returns the earliest qualifying native transfer into target.
// It is deterministic for a fixed normalized input set and makes no ownership claim.
func EarliestFunding(target string, transfers []Transfer) (model.Edge, error) {
	matches := make([]Transfer, 0)
	for _, t := range transfers {
		if t.To == target && t.From != "" && t.Lamports > 0 && t.Signature != "" {
			matches = append(matches, t)
		}
	}
	if len(matches) == 0 {
		return model.Edge{}, errors.New("no qualifying funding transfer")
	}
	sort.SliceStable(matches, func(i, j int) bool {
		if matches[i].Slot != matches[j].Slot {
			return matches[i].Slot < matches[j].Slot
		}
		return matches[i].Signature < matches[j].Signature
	})
	m := matches[0]
	return model.Edge{
		Source:       m.To,
		Target:       m.From,
		Relationship: model.FundedBy,
		Amount:       m.Lamports,
		Provenance: model.Provenance{
			Signature:        m.Signature,
			Slot:             m.Slot,
			Timestamp:        m.Timestamp,
			DerivationMethod: "earliest_qualifying_native_transfer_v0",
		},
	}, nil
}
