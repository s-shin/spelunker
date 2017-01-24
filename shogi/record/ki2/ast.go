package ki2

import "github.com/s-shin/spelunker/shogi"

type Ki2 struct {
	MetaEntries []*MetaEntry
	Notes       []*Note
	Moves       []*Move
	Result      *Result
}

type MetaEntry struct {
	Key   string
	Value string
}

type Note struct {
	Text string
}

type MovementDescriptor int

// MovementDescriptor constants.
const (
	MovementDescriptorNull MovementDescriptor = iota
	MoveUpward
	MoveDownward
	MoveVertically
	MoveFromRight
	MoveFromLeft
	MoveHorizontally
	Drop
)

type Move struct {
	Side                shogi.Side
	Destination         shogi.Position
	Same                bool
	Piece               shogi.Piece
	MovementDescriptors []MovementDescriptor
	Promotion           bool
}

type Result struct {
	Text string
}
