package shogi

import (
	"testing"
)

func TestPositionString(t *testing.T) {
	if Position11.String() != "11" {
		t.Errorf("Position#String()")
	}
}

func TestPositionOffset(t *testing.T) {
	if Position11.Offset(0, 1) != Position12 {
		t.Errorf("Position#Offset()")
	}
}

func TestPositionIsOnBoard(t *testing.T) {
	p := Position{0, 1}
	if p.IsOnBoard() {
		t.Errorf("%v should not be on board.", p)
	}
}
