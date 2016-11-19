package shogi

import "testing"

var Not = map[bool]string{true: "", false: " not "}

func TestIsCheck(t *testing.T) {
	type TestCase struct {
		board    Board
		side     Side
		expected bool
	}
	testCases := []TestCase{
		TestCase{
			board: Board{
				Position51: WOU,
				Position52: BFU,
			},
			side:     White,
			expected: true,
		},
		TestCase{
			board: Board{
				Position57: WFU,
				Position59: BOU,
			},
			side:     Black,
			expected: false,
		},
		TestCase{
			board: Board{
				Position51: WOU,
				Position55: WFU,
				Position59: BKY,
			},
			side:     White,
			expected: false,
		},
	}
	for _, c := range testCases {
		if c.board.IsCheck(c.side) != c.expected {
			t.Errorf("%s's OU should%s be judged as check:\n%s", c.side, Not[c.expected], c.board)
		}
	}
}

func TestBoardSearchMovableForBlackAboutMove(t *testing.T) {
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
		MoveTestCase{
			Board: Board{
				Position19: BOU,
				Position18: WFU,
				Position17: WKY,
			},
			From: Position19,
			To:   Position18,
			Can:  false,
		},
	}
	for _, testCase := range moveTestCases {
		movable := testCase.Board.searchMovableForBlack(testCase.From)
		if _, ok := movable[testCase.To]; ok != testCase.Can {
			t.Errorf("Piece at %v should%s be able to move to %v:\n%v", testCase.From, Not[testCase.Can], testCase.To, testCase.Board)
		}
	}
}

func TestBoardSearchDroppableForBlack(t *testing.T) {
	type DropTestCase struct {
		Board Board
		To    Position
		Piece Piece
		Can   bool
	}

	var dropTestCases = []DropTestCase{
		DropTestCase{
			Board: Board{},
			To:    Position19,
			Piece: FU,
			Can:   true,
		},
		DropTestCase{
			Board: Board{},
			To:    Position11,
			Piece: FU,
			Can:   false,
		},
		DropTestCase{
			Board: Board{},
			To:    Position51,
			Piece: KY,
			Can:   false,
		},
		DropTestCase{
			Board: Board{},
			To:    Position12,
			Piece: KE,
			Can:   false,
		},
		DropTestCase{
			Board: Board{
				Position19: BFU,
			},
			To:    Position18,
			Piece: FU,
			Can:   false,
		},
		DropTestCase{
			Board: Board{
				Position11: WOU,
				Position21: WFU,
			},
			To:    Position12,
			Piece: FU,
			Can:   true,
		},
		DropTestCase{
			Board: Board{
				Position11: WOU,
				Position21: WFU,
				Position13: BKI,
			},
			To:    Position12,
			Piece: FU,
			Can:   false,
		},
	}
	for _, testCase := range dropTestCases {
		droppable := testCase.Board.searchDroppableForBlack(testCase.Piece)
		if _, ok := droppable[testCase.To]; ok != testCase.Can {
			t.Errorf("%v should%s be able to drop to %v:\n%v", testCase.Piece, Not[testCase.Can], testCase.To, testCase.Board)
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
