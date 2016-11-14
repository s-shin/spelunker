package shogi

import "testing"

type MoveTestCase struct {
	Board Board
	From  Position
	To    Position
	Can   bool
}

var moveTestCases = []MoveTestCase{
	MoveTestCase{
		Board: Board{
			Position27: BFU,
		},
		From: Position27,
		To:   Position26,
		Can:  true,
	},
	MoveTestCase{
		Board: Board{
			Position19: BKY,
		},
		From: Position19,
		To:   Position11,
		Can:  true,
	},
	MoveTestCase{
		Board: Board{
			Position19: BKY,
			Position15: WFU,
		},
		From: Position19,
		To:   Position15,
		Can:  true,
	},
	MoveTestCase{
		Board: Board{
			Position19: BKY,
			Position15: WFU,
		},
		From: Position19,
		To:   Position11,
		Can:  false,
	},
}

func TestBoardSearchMovableForBlackAboutMove(t *testing.T) {
	not := map[bool]string{true: "", false: "not "}
	for _, testCase := range moveTestCases {
		movable := testCase.Board.searchMovableForBlack(testCase.From)
		if _, ok := movable[testCase.To]; ok != testCase.Can {
			t.Errorf("%v of %v should%s be able to move %v", testCase.From, testCase.Board, not[testCase.Can], testCase.To)
		}
	}
}

func TestBoardHas(t *testing.T) {
	b := NewBoard()
	b.SafePut(Position11, WKY)
	b.SafePut(Position12, SPN)
	if !b.Has(Position11) {
		t.Errorf("The side and piece (%s) at Position11 should be WKY", b.Get(Position11))
	}
	if b.Has(Position12) {
		t.Errorf("SideAndPieceNull should be regarded as empty.")
	}
}

func TestBoardEmptyPositions(t *testing.T) {
	b := NewBoard()
	if ps := b.EmptyPositions(); len(ps) != BoardWidth*BoardHeight {
		t.Errorf("EmptyPositions of empty board should return the slice whose length is 81 (slice size : %d).", len(ps))
	}
}
