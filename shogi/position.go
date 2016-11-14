package shogi

import (
	"fmt"
)

// Position represents a position on board.
type Position struct {
	x int
	y int
}

// Position constants.
var (
	PositionNull = Position{}
	// for i in {1..9}; do for j in {1..9}; do echo "Position$i$j = Position{$i, $j}"; done; done
	Position11 = Position{1, 1}
	Position12 = Position{1, 2}
	Position13 = Position{1, 3}
	Position14 = Position{1, 4}
	Position15 = Position{1, 5}
	Position16 = Position{1, 6}
	Position17 = Position{1, 7}
	Position18 = Position{1, 8}
	Position19 = Position{1, 9}
	Position21 = Position{2, 1}
	Position22 = Position{2, 2}
	Position23 = Position{2, 3}
	Position24 = Position{2, 4}
	Position25 = Position{2, 5}
	Position26 = Position{2, 6}
	Position27 = Position{2, 7}
	Position28 = Position{2, 8}
	Position29 = Position{2, 9}
	Position31 = Position{3, 1}
	Position32 = Position{3, 2}
	Position33 = Position{3, 3}
	Position34 = Position{3, 4}
	Position35 = Position{3, 5}
	Position36 = Position{3, 6}
	Position37 = Position{3, 7}
	Position38 = Position{3, 8}
	Position39 = Position{3, 9}
	Position41 = Position{4, 1}
	Position42 = Position{4, 2}
	Position43 = Position{4, 3}
	Position44 = Position{4, 4}
	Position45 = Position{4, 5}
	Position46 = Position{4, 6}
	Position47 = Position{4, 7}
	Position48 = Position{4, 8}
	Position49 = Position{4, 9}
	Position51 = Position{5, 1}
	Position52 = Position{5, 2}
	Position53 = Position{5, 3}
	Position54 = Position{5, 4}
	Position55 = Position{5, 5}
	Position56 = Position{5, 6}
	Position57 = Position{5, 7}
	Position58 = Position{5, 8}
	Position59 = Position{5, 9}
	Position61 = Position{6, 1}
	Position62 = Position{6, 2}
	Position63 = Position{6, 3}
	Position64 = Position{6, 4}
	Position65 = Position{6, 5}
	Position66 = Position{6, 6}
	Position67 = Position{6, 7}
	Position68 = Position{6, 8}
	Position69 = Position{6, 9}
	Position71 = Position{7, 1}
	Position72 = Position{7, 2}
	Position73 = Position{7, 3}
	Position74 = Position{7, 4}
	Position75 = Position{7, 5}
	Position76 = Position{7, 6}
	Position77 = Position{7, 7}
	Position78 = Position{7, 8}
	Position79 = Position{7, 9}
	Position81 = Position{8, 1}
	Position82 = Position{8, 2}
	Position83 = Position{8, 3}
	Position84 = Position{8, 4}
	Position85 = Position{8, 5}
	Position86 = Position{8, 6}
	Position87 = Position{8, 7}
	Position88 = Position{8, 8}
	Position89 = Position{8, 9}
	Position91 = Position{9, 1}
	Position92 = Position{9, 2}
	Position93 = Position{9, 3}
	Position94 = Position{9, 4}
	Position95 = Position{9, 5}
	Position96 = Position{9, 6}
	Position97 = Position{9, 7}
	Position98 = Position{9, 8}
	Position99 = Position{9, 9}
)

// MakePositionByTopLeftOrigin creates a position by the coordinates
// whose origin is at the top left.
// This is useful for creating a position from an array in program.
func MakePositionByTopLeftOrigin(x, y int) Position {
	return Position{BoardWidth - x, y + 1}
}

// MakePositionFromString creates a position from the string of a position.
func MakePositionFromString(s string) Position {
	var x, y int
	if n, err := fmt.Sscanf(s, "%1d%1d", &x, &y); err == nil && n == 2 {
		return Position{x, y}
	}
	fmt.Println(x, y)
	return PositionNull
}

// Reverse the position.
func (pos Position) Reverse() Position {
	return Position{BoardWidth - pos.x + 1, BoardHeight - pos.y + 1}
}

func (pos Position) String() string {
	if pos == PositionNull {
		return ""
	}
	return fmt.Sprintf("%d%d", pos.x, pos.y)
}

// IsOnBoard returns true if the position is inside board.
func (pos Position) IsOnBoard() bool {
	return pos.x >= 1 && pos.x <= BoardWidth && pos.y >= 1 && pos.y <= BoardHeight
}

// IsInPromotionZone returns true if the position is in the promotion zone of the side.
func (pos Position) IsInPromotionZone(side Side) bool {
	switch side {
	case Black:
		return pos.y <= 3
	case White:
		return pos.y >= 7
	}
	panic("Invalid side: " + side.String())
}

// Offset returns the position moved with the offset.
func (pos Position) Offset(left, down int) Position {
	p := Position{pos.x + left, pos.y + down}
	if !p.IsOnBoard() {
		return PositionNull
	}
	return p
}
