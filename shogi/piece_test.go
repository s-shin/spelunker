package shogi

import (
	"testing"
)

func TestPieceString(t *testing.T) {
	if FU.String() != "FU" {
		t.Errorf("Piece#String()")
	}
}

func TestPiecePromote(t *testing.T) {
	if !FU.CanPromote() {
		t.Error("FU should be able to be promoted.")
	}
	if FU.Promote().String() != "TO" {
		t.Error("Promoted FU should be TO.")
	}
}
