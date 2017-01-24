package shogi

import "time"

type Player struct {
	FirstName string
	LastName  string
}

type PlayerPair struct {
	Black Player
	White Player
}

type Record struct {
	StartedAt  time.Time
	Tournament string
	Location   string
	Players    PlayerPair
	Meta       map[string]string
	Moves      []Move
	Game       *Game
}

type RecordLoader interface {
	Load(filePath string) (*Record, error)
}

type RecordFormat int

// RecordType list
const (
	RecordFormatNull RecordFormat = iota
	KI2
	KIF
	CSA
)

//----------------------------

// func ReadKi2(fp io.Reader) (*Record, error) {
// 	reader := bufio.NewReader(transform.NewReader(fp, japanese.ShiftJIS.NewDecoder()))
// 	record := &Record{}
//
// 	type tokenType int
//
// 	const (
// 		TokenTypeNull tokenType = iota
// 		TokenMetaKey
// 		TokenMetaValue
// 		TokenMetaSeparator
// 		TokenMoveSide
// 		TokenMoveDestX
// 		TokenMoveDestY
// 		TokenMoveSame
// 		TokenMovePiece
// 		TokenMoveDropped
// 		TokenMoveDownward
// 		TokenMoveHorizontally
// 		TokenMoveUpward
// 		TokenMoveFromRight
// 		TokenMoveFromLeft
// 		TokenMovePromotion
// 		TokenMoveVertically
// 		TokenResult
// 		TokenComment
// 	)
//
// 	type token struct {
// 		Type tokenType
// 		Pos  int
// 		Str  string
// 	}
//
// 	const TokenNull = token{}
//
// 	type buffer struct {
// 		buf    []rune
// 		cursor int
// 	}
//
// 	scanTokenMetaKey := func(in []rune, pos int) (token, []rune) {
// 		var r []rune
// 		for _, c := range in {
// 			if unicode.IsSpace(c) {
// 				return TokenNull, in
// 			}
// 			if string(c) == "：" {
// 				break
// 			}
// 			r = append(r, c)
// 		}
// 		return token{TokenMetaKey, pos, string(r)}, in[len(r):]
// 	}
//
// 	scanTokenMetaSeparator := func(input []rune, pos int) (token, []rune) {
// 		if input[0] != "：" {
// 			return TokenNull, input
// 		}
// 		return token{TokenMetaSeparator, input[0]}, input[1:]
// 	}
//
// 	scanTokenMetaValue := func(input []rune, pos int) (token, []rune) {
//
// 	}
//
// 	game := NewGameWithStartingPositions()
// 	var moves []Move
// 	fmt.Print(game, moves)
//
// 	for {
// 		r, _, err := reader.ReadRune()
// 		if err != nil {
// 			break
// 		}
// 		switch string(r) {
// 		case "▲":
// 			if side != SideNull {
// 				side = SideNull
// 				pos = PositionNull
// 			}
// 			side = Black
// 		case "△":
// 			side = White
// 		case "１":
// 			pos.x = 1
// 		case "２":
// 			pos.x = 2
// 		case "３":
// 			pos.x = 3
// 		case "４":
// 			pos.x = 4
// 		case "５":
// 			pos.x = 5
// 		case "６":
// 			pos.x = 6
// 		case "７":
// 			pos.x = 7
// 		case "８":
// 			pos.x = 8
// 		case "９":
// 			pos.x = 9
// 		case "一":
// 			pos.y = 1
// 		case "二":
// 			pos.y = 2
// 		case "三":
// 			pos.y = 3
// 		case "四":
// 			pos.y = 4
// 		case "五":
// 			pos.y = 5
// 		case "六":
// 			pos.y = 6
// 		case "七":
// 			pos.y = 7
// 		case "八":
// 			pos.y = 8
// 		case "九":
// 			pos.y = 9
// 		}
// 	}
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 	again:
// 		switch i {
// 		case Begin:
// 			if strings.HasPrefix(line, "▲") {
// 				i = Body
// 				goto again
// 			}
// 		case Body:
// 			if strings.HasPrefix(line, "まで") {
// 				i = End
// 				goto again
// 			}
// 			characters := []rune(line)
// 			side := SideNull
// 			pos := PositionNull
// 			for _, c := range characters {
// 				switch string(c) {
// 				case "▲":
// 					if side != SideNull {
// 						side = SideNull
// 						pos = PositionNull
// 					}
// 					side = Black
// 				case "△":
// 					side = White
// 				case "１":
// 					pos.x = 1
// 				case "２":
// 					pos.x = 2
// 				case "３":
// 					pos.x = 3
// 				case "４":
// 					pos.x = 4
// 				case "５":
// 					pos.x = 5
// 				case "６":
// 					pos.x = 6
// 				case "７":
// 					pos.x = 7
// 				case "８":
// 					pos.x = 8
// 				case "９":
// 					pos.x = 9
// 				case "一":
// 					pos.y = 1
// 				case "二":
// 					pos.y = 2
// 				case "三":
// 					pos.y = 3
// 				case "四":
// 					pos.y = 4
// 				case "五":
// 					pos.y = 5
// 				case "六":
// 					pos.y = 6
// 				case "七":
// 					pos.y = 7
// 				case "八":
// 					pos.y = 8
// 				case "九":
// 					pos.y = 9
// 				}
// 			}
// 		case End:
// 		}
// 	}
// 	return record, nil
// }
