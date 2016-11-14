package shogi

import "testing"

type test struct {
	t    *testing.T
	Game *Game
}

func (t *test) move(from Position, to Position, promote bool) {
	if _, err := t.Game.Move(from, to, promote); err != nil {
		t.t.Fatalf(
			"It should be possible to move from %s to %s (error: %s).\n%s",
			from.String(),
			to.String(),
			err.Error(),
			t.Game.AppliedGame().String(),
		)
	}
}

func (t *test) drop(piece Piece, to Position) {
	if _, err := t.Game.Drop(piece, to); err != nil {
		t.t.Fatalf(
			"It should be possible to drop %s to %s (error: %s).\n%s",
			piece.String(),
			to.String(),
			err.Error(),
			t.Game.AppliedGame().String(),
		)
	}
}

func TestGame(t *testing.T) {
	test := &test{
		t:    t,
		Game: NewGameWithStartingPositions(),
	}
	test.move(Position77, Position76, false)
	test.move(Position33, Position34, false)
	test.move(Position88, Position22, true)
	test.move(Position31, Position22, false)
	test.move(Position89, Position77, false)
	test.drop(KA, Position55)
	test.move(Position77, Position65, false)
	test.move(Position55, Position99, true)
	test.move(Position65, Position53, false)
	test.move(Position99, Position98, false)
	test.move(Position53, Position41, true)
}
