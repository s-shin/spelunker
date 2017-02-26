package shogi

import "errors"

type Game struct {
	record *Record
}

func NewGameWithBoard(board Board) *Game {
	r := &Record{
		StartingSetup: &State{
			Board:     board,
			BlackHand: Hand{},
			WhiteHand: Hand{},
			NextTurn:  Black,
		},
		Actions: []Action{},
	}
	return &Game{r}
}

// Move appends a move to the record in this game.
// If the move is against the rules, an error will be returned.
func (g *Game) Move(from, to Position, promote bool) error {
	state, err := g.State()
	if err != nil {
		return err
	}
	piece := state.Board.Get(from).Piece
	if piece == PieceNull {
		return errors.New("foo")
	}
	if promote {
		if !piece.CanPromote() {
			return errors.New("foo2")
		}
		piece = piece.Promote()
	}
	move := &Move{
		Side:  state.NextTurn,
		Piece: piece,
		From:  from,
		To:    to,
	}
	if err := state.Board.VerifyMove(move); err != nil {
		return err
	}
	g.AppendAction(move)
	return nil
}

// Drop appends a move as drop to the record in this game.
// If the move is against the rules, an error will be returned.
func (g *Game) Drop(piece Piece, to Position) error {
	state, err := g.State()
	if err != nil {
		return err
	}
	move := &Move{
		Side:  state.NextTurn,
		Piece: piece,
		To:    to,
	}
	if err := state.Board.VerifyMove(move); err != nil {
		return err
	}
	g.AppendAction(move)
	return nil
}

// AppendAction appends any action to the record in this game.
// Unlike Move and Drop, AppendAction doesn't check the rules.
func (g *Game) AppendAction(a Action) {
	g.record.Actions = append(g.record.Actions, a)
}

// Record returns current record of this game.
func (g *Game) Record() *Record {
	return g.record
}

// State return the current state.
func (g *Game) State() (*State, error) {
	return g.record.GetLatestState()
}
