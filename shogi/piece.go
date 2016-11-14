package shogi

// Piece is "koma" in Japanese.
type Piece int

// Piece constants.
const (
	PieceNull Piece = iota
	OU
	HI
	RY
	KA
	UM
	KI
	GI
	NG
	KE
	NK
	KY
	NY
	FU
	TO
)

type pieceStrings struct {
	CSA    string
	USI    string
	En     string
	EnAlt  string
	EnAbbr string
	Ja     string
	JaAlt  string
	JaAbbr string
}

var pieceStringsMap = map[Piece]pieceStrings{
	PieceNull: pieceStrings{
		CSA: "* ",
	},
	OU: pieceStrings{
		CSA:   "OU",
		USI:   "K",
		En:    "King",
		Ja:    "王",
		JaAlt: "玉",
	},
	HI: pieceStrings{
		CSA: "HI",
		USI: "R",
		En:  "Rook",
		Ja:  "飛",
	},
	RY: pieceStrings{
		CSA:   "RY",
		USI:   "+R",
		En:    "Promoted rook",
		EnAlt: "Dragon",
		Ja:    "龍",
	},
	KA: pieceStrings{
		CSA: "KA",
		USI: "B",
		En:  "Bishop",
		Ja:  "角",
	},
	UM: pieceStrings{
		CSA:   "UM",
		USI:   "+B",
		En:    "Promoted bishop",
		EnAlt: "Horse",
		Ja:    "馬",
	},
	KI: pieceStrings{
		CSA: "KI",
		USI: "G",
		En:  "Gold",
		Ja:  "金",
	},
	GI: pieceStrings{
		CSA: "GI",
		USI: "S",
		En:  "Silver",
		Ja:  "銀",
	},
	NG: pieceStrings{
		CSA:    "NG",
		USI:    "+S",
		En:     "Promoted silver",
		Ja:     "成銀",
		JaAbbr: "全",
	},
	KE: pieceStrings{
		CSA: "KE",
		USI: "N",
		En:  "Knight",
		Ja:  "桂",
	},
	NK: pieceStrings{
		CSA:    "NK",
		USI:    "+N",
		En:     "Promoted knight",
		Ja:     "成桂",
		JaAbbr: "圭",
	},
	KY: pieceStrings{
		CSA: "KY",
		USI: "L",
		En:  "Lance",
		Ja:  "香",
	},
	NY: pieceStrings{
		CSA:    "NY",
		USI:    "+L",
		En:     "Promoted pawn",
		Ja:     "成香",
		JaAbbr: "杏",
	},
	FU: pieceStrings{
		CSA: "FU",
		USI: "P",
		En:  "Pawn",
		Ja:  "歩",
	},
	TO: pieceStrings{
		CSA: "TO",
		USI: "+P",
		En:  "Promoted pawn",
		Ja:  "と",
	},
}

var promoteMap = map[Piece]Piece{
	FU: TO,
	KY: NY,
	KE: NK,
	GI: NG,
	KA: UM,
	HI: RY,
}

func (p Piece) String() string {
	return pieceStringsMap[p].CSA
}

// MakePieceFromString creates a piece from the string of a piece.
func MakePieceFromString(s string) Piece {
	for piece, pieceStrings := range pieceStringsMap {
		if s == pieceStrings.CSA {
			return piece
		}
	}
	return PieceNull
}

// CanPromote returns true if the piece can promote.
func (p Piece) CanPromote() bool {
	_, ok := promoteMap[p]
	return ok
}

// Promote the piece.
func (p Piece) Promote() Piece {
	if p.CanPromote() {
		return promoteMap[p]
	}
	return PieceNull
}

// Demote the piece.
func (p Piece) Demote() Piece {
	for piece, promoted := range promoteMap {
		if p == promoted {
			return piece
		}
	}
	return PieceNull
}

// ForceDemote forcely demotes the piece.
func (p Piece) ForceDemote() Piece {
	if p.CanPromote() {
		return p
	}
	return p.Demote()
}
