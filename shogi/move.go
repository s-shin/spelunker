package shogi

// Move represents one action of players, also a drop action.
type Move struct {
	Side  Side
	Piece Piece
	From  Position
	To    Position
}

// IsDrop returns true if the move is drop action.
func (m *Move) IsDrop() bool {
	return m.From == PositionNull
}

// SideAndPiece of the side and the piece of the move.
func (m *Move) SideAndPiece() SideAndPiece {
	return SideAndPiece{Side: m.Side, Piece: m.Piece}
}

func (m *Move) String() string {
	return m.Side.String() + m.From.String() + m.To.String() + m.Piece.String()
}

// ActionType of Action interface.
func (m *Move) ActionType() string {
	return "move"
}
