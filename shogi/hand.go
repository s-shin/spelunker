package shogi

import "strings"

// Hand is "komadai" in Japanese.
type Hand map[Piece]int

func (hand Hand) String() string {
	size := 0
	for _, num := range hand {
		size += 2 * num
	}
	buf := make([]string, 0, size)
	for piece := range hand {
		buf = append(buf, "00", piece.String())
	}
	return strings.Join(buf, "")
}

// Clone the hand.
func (hand Hand) Clone() Hand {
	newHand := make(Hand)
	for k, v := range hand {
		newHand[k] = v
	}
	return newHand
}

// Capture a piece into the hand.
func (hand Hand) Capture(piece Piece) {
	hand[piece.ForceDemote()]++
}

// Drop a piece from the hand.
func (hand Hand) Drop(piece Piece) {
	hand[piece.ForceDemote()]--
}
