package shogi

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	g := NewGameWithBoard(NewHirateBoard())

	type move struct {
		From    Position
		To      Position
		Promote bool
	}
	type drop struct {
		Piece Piece
		To    Position
	}
	cases := []interface{}{
		move{Position77, Position76, false},
		move{Position33, Position34, false},
		move{Position88, Position22, true},
		move{Position31, Position22, false},
		move{Position89, Position77, false},
		drop{KA, Position55},
		move{Position77, Position65, false},
		move{Position55, Position99, true},
		move{Position65, Position53, false},
		move{Position99, Position98, false},
		move{Position53, Position41, true},
	}
	for _, c := range cases {
		switch x := c.(type) {
		case move:
			if err := g.Move(x.From, x.To, x.Promote); err != nil {
				t.Fatal(err, x)
			}
		case drop:
			if err := g.Drop(x.Piece, x.To); err != nil {
				t.Fatal(err, x)
			}
		}
	}
	state, err := g.State()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(state.String())
}
