package shogi

import (
	"testing"
)

func TestSideString(t *testing.T) {
	if Black.String() != "+" {
		t.Errorf("Side#String()")
	}
}
