package shogi

// SideAndPiece represents a piece of a side on board.
type SideAndPiece struct {
	Side  Side
	Piece Piece
}

// SideAndPiece constants.
var (
	SideAndPieceNull = SideAndPiece{}
	SPN              = SideAndPieceNull
	BOU              = SideAndPiece{Black, OU}
	BHI              = SideAndPiece{Black, HI}
	BRY              = SideAndPiece{Black, RY}
	BKA              = SideAndPiece{Black, KA}
	BUM              = SideAndPiece{Black, UM}
	BKI              = SideAndPiece{Black, KI}
	BGI              = SideAndPiece{Black, GI}
	BNG              = SideAndPiece{Black, NG}
	BKE              = SideAndPiece{Black, KE}
	BNK              = SideAndPiece{Black, NK}
	BKY              = SideAndPiece{Black, KY}
	BNY              = SideAndPiece{Black, NY}
	BFU              = SideAndPiece{Black, FU}
	BTO              = SideAndPiece{Black, TO}
	WOU              = SideAndPiece{White, OU}
	WHI              = SideAndPiece{White, HI}
	WRY              = SideAndPiece{White, RY}
	WKA              = SideAndPiece{White, KA}
	WUM              = SideAndPiece{White, UM}
	WKI              = SideAndPiece{White, KI}
	WGI              = SideAndPiece{White, GI}
	WNG              = SideAndPiece{White, NG}
	WKE              = SideAndPiece{White, KE}
	WNK              = SideAndPiece{White, NK}
	WKY              = SideAndPiece{White, KY}
	WNY              = SideAndPiece{White, NY}
	WFU              = SideAndPiece{White, FU}
	WTO              = SideAndPiece{White, TO}
)

func (sap SideAndPiece) String() string {
	return sap.Side.String() + sap.Piece.String()
}

// IsOnInvalidPosition returns true if a SideAndPiece on the position
// can be moved to at leaset one position.
func (sap SideAndPiece) IsOnInvalidPosition(pos Position) bool {
	if sap.Side == White {
		pos = pos.Reverse()
	}
	switch sap.Piece {
	case FU, KY:
		return pos.Y == 1
	case KE:
		return pos.Y == 2
	}
	return false
}

// Reverse the SideAndPiece.
func (sap SideAndPiece) Reverse() SideAndPiece {
	return SideAndPiece{
		Side:  sap.Side.Reverse(),
		Piece: sap.Piece,
	}
}
