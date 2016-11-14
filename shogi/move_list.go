package shogi

import "strings"

// MoveList is list of moves.
type MoveList []Move

// Clone the move list.
func (ml MoveList) Clone() MoveList {
	t := make(MoveList, len(ml))
	copy(t, ml)
	return t
}

func (ml MoveList) String() string {
	strs := make([]string, len(ml))
	for i, move := range ml {
		strs[i] = move.String()
	}
	return strings.Join(strs, "\n")
}
