package shogi

import (
	"testing"
)

func TestMoveString(t *testing.T) {
	var m Move
	m = Move{Side: Black, Piece: TO, From: Position11, To: Position12}
	if m.String() != "+1112TO" {
		t.Errorf("Move#String()")
	}
	m = Move{Side: White, Piece: KA, To: Position55}
	if m.String() != "-55KA" {
		t.Errorf("Move#String()")
	}
}
