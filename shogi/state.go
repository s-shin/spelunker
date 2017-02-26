package shogi

import "strings"

// State is a state of a game.
type State struct {
	Board     Board
	BlackHand Hand
	WhiteHand Hand
	NextTurn  Side
}

func (s *State) String() string {
	strs := []string{
		"'Board:",
		s.Board.String(),
		"'Black's hand:",
		"P" + Black.String() + s.BlackHand.String(),
		"'White's hand:",
		"P" + White.String() + s.WhiteHand.String(),
		"'Next turn:",
		s.NextTurn.String(),
	}
	return strings.Join(strs, "\n")
}

func (s *State) ApplyMove(m *Move) error {
	captured, err := s.Board.ApplyMove(m)
	if err != nil {
		return err
	}
	if captured != PieceNull {
		s.GetHand(s.NextTurn).Capture(captured)
	}
	if m.IsDrop() {
		s.GetHand(s.NextTurn).Drop(m.Piece)
	}
	s.NextTurn = s.NextTurn.Reverse()
	return nil
}

func (s *State) GetHand(side Side) Hand {
	if side == Black {
		return s.BlackHand
	}
	return s.WhiteHand
}

func (s *State) Clone() *State {
	return &State{
		Board:     s.Board.Clone(),
		BlackHand: s.BlackHand.Clone(),
		WhiteHand: s.WhiteHand.Clone(),
		NextTurn:  s.NextTurn,
	}
}
