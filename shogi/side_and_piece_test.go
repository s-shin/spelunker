package shogi

import (
	"testing"
)

func TestSideAndPieceString(t *testing.T) {
	if BFU.String() != "+FU" {
		t.Errorf("SideAndPiece#String()")
	}
}
