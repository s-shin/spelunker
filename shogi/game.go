package shogi

import "strings"

// Game contains all information to construct one game.
type Game struct {
	Start       Board
	Moves       MoveList
	HandOfBlack Hand
	HandOfWhite Hand
	Turn        Side
}

// NewGame creates a new empty game.
func NewGame() *Game {
	return &Game{
		Start:       NewBoard(),
		Moves:       make(MoveList, 0),
		HandOfBlack: make(Hand),
		HandOfWhite: make(Hand),
		Turn:        SideNull,
	}
}

// NewGameWithStartingPositions creates a new game being set up with the default staring positions (hirate).
func NewGameWithStartingPositions() *Game {
	game := NewGame()
	game.Turn = Black
	for i, sap := range []SideAndPiece{
		WKY, WKE, WGI, WKI, WOU, WKI, WGI, WKE, WKY,
		SPN, WHI, SPN, SPN, SPN, SPN, SPN, WKA, SPN,
		WFU, WFU, WFU, WFU, WFU, WFU, WFU, WFU, WFU,
		SPN, SPN, SPN, SPN, SPN, SPN, SPN, SPN, SPN,
		SPN, SPN, SPN, SPN, SPN, SPN, SPN, SPN, SPN,
		SPN, SPN, SPN, SPN, SPN, SPN, SPN, SPN, SPN,
		BFU, BFU, BFU, BFU, BFU, BFU, BFU, BFU, BFU,
		SPN, BKA, SPN, SPN, SPN, SPN, SPN, BHI, SPN,
		BKY, BKE, BGI, BKI, BOU, BKI, BGI, BKE, BKY,
	} {
		if sap != SideAndPieceNull {
			pos := MakePositionByTopLeftOrigin(i%BoardWidth, i/BoardWidth)
			game.Start.SafePut(pos, sap)
		}
	}
	return game
}

// Clone the game.
func (game *Game) Clone() *Game {
	return &Game{
		Start:       game.Start.Clone(),
		Moves:       game.Moves.Clone(),
		HandOfBlack: game.HandOfBlack.Clone(),
		HandOfWhite: game.HandOfWhite.Clone(),
		Turn:        game.Turn,
	}
}

func (game *Game) String() string {
	strs := []string{
		"'Board:",
		game.Start.String(),
		"'Black's hand:",
		"P" + Black.String() + game.HandOfBlack.String(),
		"'White's hand:",
		"P" + White.String() + game.HandOfWhite.String(),
		"'Turn:",
		game.Turn.String(),
	}
	if len(game.Moves) > 0 {
		strs = append(strs, []string{
			"'Moves:",
			game.Moves.String(),
		}...)
	}
	return strings.Join(strs, "\n")
}

// GetHand returns the hand of the side.
func (game *Game) GetHand(side Side) Hand {
	switch side {
	case Black:
		return game.HandOfBlack
	case White:
		return game.HandOfWhite
	}
	panic("error")
}

// ChangeTurn toggles the turn.
func (game *Game) ChangeTurn() {
	if game.Turn == Black {
		game.Turn = White
	} else {
		game.Turn = Black
	}
}

// ApplyOneMove applies the head of the move list and removes it.
func (game *Game) ApplyOneMove() error {
	move := game.Moves[0]
	game.Moves = game.Moves[1:]
	if piece, err := game.Start.ApplyMove(move); err != nil {
		return err
	} else if piece != PieceNull {
		game.GetHand(game.Turn).Capture(piece)
	}
	if move.IsDrop() {
		game.GetHand(game.Turn).Drop(move.Piece)
	}
	game.ChangeTurn()
	return nil
}

// ApplyMoves applies the given number of moves from the head of the move list.
func (game *Game) ApplyMoves(n int) error {
	for i := 0; i < n; i++ {
		if err := game.ApplyOneMove(); err != nil {
			return err
		}
	}
	return nil
}

// ApplyAllMoves applies all moves in the move list.
func (game *Game) ApplyAllMoves() error {
	return game.ApplyMoves(len(game.Moves))
}

// AppliedGame returns the game applied all moves to.
func (game *Game) AppliedGame() *Game {
	g := game.Clone()
	g.ApplyAllMoves()
	return g
}

// Move makes a move action easy and returns the applied game with error.
func (game *Game) Move(from, to Position, promote bool) (*Game, error) {
	g := game.AppliedGame()
	piece := g.Start.Get(from).Piece
	if promote {
		piece = piece.Promote()
	}
	move := Move{
		Side:  g.Turn,
		Piece: piece,
		From:  from,
		To:    to,
	}
	g.Moves = append(g.Moves, move)
	if err := g.ApplyAllMoves(); err != nil {
		return nil, err
	}
	game.Moves = append(game.Moves, move)
	return g, nil
}

// Drop makes a drop action easy and returns the applied game with error.
func (game *Game) Drop(piece Piece, to Position) (*Game, error) {
	g := game.AppliedGame()
	move := Move{
		Side:  g.Turn,
		Piece: piece,
		To:    to,
	}
	g.Moves = append(g.Moves, move)
	if err := g.ApplyAllMoves(); err != nil {
		return nil, err
	}
	game.Moves = append(game.Moves, move)
	return g, nil
}
