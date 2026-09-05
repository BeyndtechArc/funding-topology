package extract

import (
	"testing"
	"time"

	"github.com/BeyndtechArc/funding-topology/internal/model"
)

func TestFeePayerEdge(t *testing.T) {
	tx := Transaction{Signature: "sig-1", Slot: 42, Timestamp: time.Unix(1_700_000_000, 0).UTC(), FeePayer: "payer", Signers: []string{"seed"}}
	edge, err := FeePayerEdge("seed", tx)
	if err != nil {
		t.Fatal(err)
	}
	if edge.Relationship != model.FeePaidBy || edge.Source != "seed" || edge.Target != "payer" {
		t.Fatalf("unexpected edge: %+v", edge)
	}
}
