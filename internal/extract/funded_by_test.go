package extract

import (
	"testing"
	"time"

	"github.com/BeyndtechArc/funding-topology/internal/model"
)

func TestEarliestFunding(t *testing.T) {
	now := time.Unix(1_700_000_000, 0).UTC()
	transfers := []Transfer{
		{From: "later-funder", To: "seed", Lamports: 20, Signature: "sig-b", Slot: 101, Timestamp: now.Add(time.Second)},
		{From: "first-funder", To: "seed", Lamports: 10, Signature: "sig-a", Slot: 100, Timestamp: now},
	}
	edge, err := EarliestFunding("seed", transfers)
	if err != nil {
		t.Fatal(err)
	}
	if edge.Relationship != model.FundedBy {
		t.Fatalf("relationship=%s", edge.Relationship)
	}
	if edge.Source != "seed" || edge.Target != "first-funder" {
		t.Fatalf("unexpected edge: %+v", edge)
	}
	if edge.Provenance.Signature != "sig-a" || edge.Amount != 10 {
		t.Fatalf("unexpected provenance: %+v", edge)
	}
}

func TestEarliestFundingRejectsEmpty(t *testing.T) {
	if _, err := EarliestFunding("seed", nil); err == nil {
		t.Fatal("expected error")
	}
}
