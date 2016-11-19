package shogi

import (
	"fmt"
	"strings"
)

// Board size constants.
const (
	BoardWidth  = 9
	BoardHeight = 9
)

// Board is "ban" in Japanese.
type Board map[Position]SideAndPiece

// NewBoard creates a new board.
func NewBoard() Board {
	return make(Board)
}

// SafePut puts a piece to the board safely.
func (b Board) SafePut(pos Position, sap SideAndPiece) {
	if b.Has(pos) {
		panic(fmt.Sprintf("%s must be empty.", pos.String()))
	}
	b[pos] = sap
}

// SafeRemove removes a piece from the board safely.
func (b Board) SafeRemove(pos Position) {
	if !b.Has(pos) {
		panic(fmt.Sprintf("%s must not be empty.", pos.String()))
	}
	delete(b, pos)
}

// Get the SideAndPiece at the position.
func (b Board) Get(pos Position) SideAndPiece {
	if sap, ok := b[pos]; ok {
		return sap
	}
	return SideAndPieceNull
}

// Has return true if the board has a piece at the position.
func (b Board) Has(pos Position) bool {
	return b.Get(pos) != SideAndPieceNull
}

// Clone the board.
func (b Board) Clone() Board {
	newBoard := make(Board)
	for k, v := range b {
		newBoard[k] = v
	}
	return newBoard
}

func (b Board) String() string {
	rows := make([]string, 0, BoardHeight)
	for y := 1; y <= BoardHeight; y++ {
		row := make([]string, 0, BoardWidth+1)
		row = append(row, fmt.Sprintf("P%d", y))
		for x := BoardWidth; x >= 1; x-- {
			row = append(row, b.Get(Position{x, y}).String())
		}
		rows = append(rows, strings.Join(row, ""))
	}
	return strings.Join(rows, "\n")
}

// Reverse the board.
func (b Board) Reverse() Board {
	newBoard := Board{}
	for pos, sap := range b {
		newBoard[pos.Reverse()] = sap.Reverse()
	}
	return newBoard
}

// EmptyPositions returns the positions not being empty.
func (b Board) EmptyPositions() []Position {
	ps := make([]Position, 0, BoardWidth*BoardHeight-len(b))
	for x := 1; x <= BoardWidth; x++ {
		for y := 1; y <= BoardHeight; y++ {
			p := Position{x, y}
			if !b.Has(p) {
				ps = append(ps, p)
			}
		}
	}
	return ps
}

// Each iterates each non-empty SideAndPiece on board with the fn.
// The iterates can be broken by returning false from the fn.
func (b Board) Each(fn func(pos Position, sap SideAndPiece) bool) {
	for pos, sap := range b {
		if sap == SideAndPieceNull {
			continue
		}
		if !fn(pos, sap) {
			break
		}
	}
}

// Search returns the positions matching the cond.
func (b Board) Search(cond func(pos Position, sap SideAndPiece) bool) []Position {
	var results []Position
	b.Each(func(pos Position, sap SideAndPiece) bool {
		if cond(pos, sap) {
			results = append(results, pos)
		}
		return true
	})
	return results
}

// ApplyMove applies a move action.
func (b Board) ApplyMove(m Move) (Piece, error) {
	if !b.CanApplyMove(m) {
		return PieceNull, NewMoveError(m)
	}
	if m.IsDrop() {
		b[m.To] = m.SideAndPiece()
	} else {
		fromSap := b.Get(m.From)
		toSap := b.Get(m.To)
		if fromSap.Piece != m.Piece {
			p := fromSap.Piece.Promote()
			if p != m.Piece {
				// FIXME: should return better error
				return PieceNull, NewInvalidStateError("fromSap is not correspond to m")
			}
			fromSap.Piece = p
		}
		b[m.To] = fromSap
		b.SafeRemove(m.From)
		if toSap != SideAndPieceNull {
			return toSap.Piece, nil
		}
	}
	return PieceNull, nil
}

// CanApplyMove returns true if the move can be applied.
func (b Board) CanApplyMove(m Move) bool {
	if m.IsDrop() {
		if b.Has(m.To) {
			return false
		}
	} else {
		if !b.Has(m.From) {
			return false
		}
	}
	var moves map[Move]struct{}
	if m.IsDrop() {
		moves = b.SearchDrops(m.SideAndPiece())
	} else {
		moves = b.SearchMoves(m.From)
	}
	_, ok := moves[m]
	return ok
}

// SearchMoves return available moves of the SideAndMove at the position.
func (b Board) SearchMoves(pos Position) map[Move]struct{} {
	movable := b.searchMovable(pos)
	moves := make(map[Move]struct{})
	for to := range movable {
		for piece := range movable[to] {
			moves[Move{Side: b[pos].Side, Piece: piece, From: pos, To: to}] = struct{}{}
		}
	}
	return moves
}

func (b Board) searchMovable(pos Position) map[Position]map[Piece]struct{} {
	switch b[pos].Side {
	case Black:
		return b.searchMovableForBlack(pos)
	case White:
		return b.searchMovableForWhite(pos)
	default:
		panic("Invalid position: " + pos.String())
	}
}

func (b Board) searchMovableForBlack(pos Position) map[Position]map[Piece]struct{} {
	piece := b[pos].Piece
	movable := make(map[Position]map[Piece]struct{})
	canMoveTo := func(p Position) bool {
		sap, ok := b[p]
		return !ok || sap.Side != Black
	}
	setMovable := func(pos Position) {
		if _, ok := movable[pos]; ok {
			return
		}
		availablePieces := make(map[Piece]struct{})
		sap := SideAndPiece{Black, piece}
		if sap.IsOnInvalidPosition(pos) {
			// promote forcely
			availablePieces[piece.Promote()] = struct{}{}
		} else {
			availablePieces[piece] = struct{}{}
			if pos.IsInPromotionZone(Black) {
				availablePieces[piece.Promote()] = struct{}{}
			}
		}
		movable[pos] = availablePieces
	}
	setMovableByLine := func(dx, dy int) {
		for p := pos.Offset(dx, dy); p.IsOnBoard() && canMoveTo(p); p = p.Offset(dx, dy) {
			setMovable(p)
			if b[p].Side == White {
				break // get white's piece and stop
			}
		}
	}
	setMovableWithOffsets := func(offsets [][2]int) {
		for _, offset := range offsets {
			p := pos.Offset(offset[0], offset[1])
			if p.IsOnBoard() && canMoveTo(p) {
				setMovable(p)
			}
		}
	}
	switch piece {
	case FU:
		setMovableWithOffsets([][2]int{{0, -1}})
	case KY:
		setMovableByLine(0, -1)
	case KE:
		setMovableWithOffsets([][2]int{{-1, -2}, {1, -2}})
	case GI:
		setMovableWithOffsets([][2]int{
			{1, -1}, {0, -1}, {-1, -1},
			/* .. */ /* .. */ /* .. */
			{1, 1} /* .. */, {-1, 1},
		})
	case KI, TO, NY, NK, NG:
		setMovableWithOffsets([][2]int{
			{1, -1}, {0, -1}, {-1, -1},
			{1, 0} /* .. */, {0, -1},
			/* .. */ {0, 1}, /* .. */
		})
	case KA, UM:
		setMovableByLine(1, 1)
		setMovableByLine(1, -1)
		setMovableByLine(-1, 1)
		setMovableByLine(-1, -1)
		if piece == UM {
			setMovableWithOffsets([][2]int{
				/* .. */ {0, -1}, /* .. */
				{1, 0} /* .. */, {0, -1},
				/* .. */ {0, 1}, /* .. */
			})
		}
	case HI, RY:
		setMovableByLine(1, 0)
		setMovableByLine(-1, 0)
		setMovableByLine(0, 1)
		setMovableByLine(0, -1)
		if piece == RY {
			setMovableWithOffsets([][2]int{
				{1, -1} /* .. */, {-1, -1},
				/* .. */ /* .. */ /* .. */
				{1, 1} /* .. */, {-1, 1},
			})
		}
	case OU:
		setMovableWithOffsets([][2]int{
			{1, -1}, {0, -1}, {-1, -1},
			{1, 0} /* .. */, {0, -1},
			{1, 1}, {0, 1}, {-1, 1},
		})
		for mp := range movable {
			bb := b.Clone()
			bb.SafeRemove(pos)
			bb[mp] = SideAndPiece{Side: Black, Piece: OU}
			if bb.IsCheck(Black) {
				delete(movable, mp)
			}
		}
	}
	return movable
}

func (b Board) searchMovableForWhite(pos Position) map[Position]map[Piece]struct{} {
	movable := make(map[Position]map[Piece]struct{})
	for p, v := range b.Reverse().searchMovableForBlack(pos.Reverse()) {
		movable[p.Reverse()] = v
	}
	return movable
}

// SearchDrops return available drops, each of which a Move instance, of the SideAndMove at the position.
func (b Board) SearchDrops(sap SideAndPiece) map[Move]struct{} {
	var droppable map[Position]struct{}
	switch sap.Side {
	case Black:
		droppable = b.searchDroppableForBlack(sap.Piece)
	case White:
		droppable = b.searchDroppableForWhite(sap.Piece)
	default:
		// ERROR
	}
	drops := make(map[Move]struct{})
	for p := range droppable {
		drops[Move{Side: sap.Side, Piece: sap.Piece, To: p}] = struct{}{}
	}
	return drops
}

func (b Board) searchDroppableForBlack(piece Piece) map[Position]struct{} {
	droppable := make(map[Position]struct{})
	var fuCols map[int]struct{} // columns containing black's fu. lazily initialized.
	for _, p := range b.EmptyPositions() {
		switch piece {
		case FU:
			if p.y == 1 {
				continue
			}
			// check nifu
			if fuCols == nil {
				fuCols = make(map[int]struct{})
				for x := 1; x <= BoardWidth; x++ {
					for y := 1; y <= BoardHeight; y++ {
						if sap2, ok := b[Position{x, y}]; ok {
							if sap2.Piece == FU && sap2.Side == Black {
								fuCols[x] = struct{}{}
								break
							}
						}
					}
				}
			}
			if _, ok := fuCols[p.x]; ok {
				continue
			}
			// check uchifuzume
			if sap := b.Get(Position{p.x, p.y - 1}); sap == WOU {
				bb := b.Clone()
				bb[p] = BFU
				if bb.IsCheckmate(White) {
					continue
				}
			}
		case KY:
			if p.y == 1 {
				continue
			}
		case KE:
			if p.y <= 2 {
				continue
			}
		}
		droppable[p] = struct{}{}
	}
	return droppable
}

func (b Board) searchDroppableForWhite(piece Piece) map[Position]struct{} {
	droppable := make(map[Position]struct{})
	for p := range b.Reverse().searchDroppableForBlack(piece) {
		droppable[p.Reverse()] = struct{}{}
	}
	return droppable
}

// IsCheck judges OU of the side is check.
func (b Board) IsCheck(side Side) bool {
	ouPos := PositionNull
	if ouPositions := b.Search(func(pos Position, sap SideAndPiece) bool {
		return sap.Side == side && sap.Piece == OU
	}); len(ouPositions) > 0 {
		ouPos = ouPositions[0]
	}
	if ouPos == PositionNull {
		panic("OU not found in " + side.String() + " side.")
	}
	isCheck := false
	b.Each(func(pos Position, sap SideAndPiece) bool {
		if sap.Side == side {
			return true
		}
		movable := b.searchMovable(pos)
		for p := range movable {
			if p == ouPos {
				isCheck = true
				return false
			}
		}
		return true
	})
	return isCheck
}

// IsCheckmate judges OU of the side is checkmate.
func (b Board) IsCheckmate(side Side) bool {
	ouPos := PositionNull
	if ouPositions := b.Search(func(pos Position, sap SideAndPiece) bool {
		return sap.Side == side && sap.Piece == OU
	}); len(ouPositions) > 0 {
		ouPos = ouPositions[0]
	}
	if ouPos == PositionNull {
		panic("OU not found in " + side.String() + " side.")
	}
	ouMovable := b.searchMovable(ouPos)
	for pos := range ouMovable {
		bb := b.Clone()
		bb[pos] = SideAndPiece{Side: side, Piece: OU}
		if !bb.IsCheck(side) {
			return false
		}
	}
	return true
}
