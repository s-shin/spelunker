package ki2

import (
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
	"github.com/s-shin/spelunker/parsec"
	"github.com/s-shin/spelunker/shogi"
	"github.com/s-shin/spelunker/shogi/record/recordutil"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func Load(filePath string) (*shogi.Record, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(transform.NewReader(fp, transform.Chain(japanese.ShiftJIS.NewDecoder(), recordutil.NewNewlineNormalizer())))
	if err != nil {
		errors.Wrapf(err, "failed to read the file: %s", filePath)
	}
	reader := parsec.NewInMemorySourceReader(string(data))
	ki2, err := parse(reader)
	if err != nil {
		return nil, err
	}
	pp.Println(ki2)
	return nil, nil
}

var mbnum = map[rune]int{
	'１': 1, '２': 2, '３': 3, '４': 4, '５': 5, '６': 6, '７': 7, '８': 8, '９': 9,
	'一': 1, '二': 2, '三': 3, '四': 4, '五': 5, '六': 6, '七': 7, '八': 8, '九': 9,
}

var rune2mvdesc = map[rune]MovementDescriptor{
	'打': Drop,
	'上': MoveUpward, '引': MoveDownward, '寄': MoveHorizontally,
	'右': MoveFromRight, '左': MoveFromLeft, '直': MoveVertically,
}

func parse(r parsec.SourceReader) (*Ki2, error) {
	ws := parsec.RuneIn(" 　")
	ws0 := parsec.Many(ws)

	blankLine := parsec.And(ws0, parsec.Rune('\n'))

	metaEntry := parsec.MapValueAsResults(
		parsec.And(
			parsec.StrByRuneNotIn("*：\n", 0),
			parsec.Rune('：'),
			parsec.StrByRuneNotIn("\n", 0),
			parsec.Rune('\n'),
		),
		func(rs []parsec.ParseResult) (interface{}, error) {
			entry := &MetaEntry{
				Key:   rs[0].Value().(string),
				Value: rs[2].Value().(string),
			}
			return entry, nil
		},
	)

	metaEntries := parsec.MapValueAsResults(
		parsec.ManyWithoutNil(
			parsec.Or(
				metaEntry,
				parsec.ToNil(blankLine),
			),
		),
		func(rs []parsec.ParseResult) (interface{}, error) {
			entries := make([]*MetaEntry, 0, len(rs))
			for _, ret := range rs {
				entries = append(entries, ret.Value().(*MetaEntry))
			}
			return entries, nil
		},
	)

	note := parsec.Map(
		parsec.StrByAnd(
			parsec.Rune('*'),
			parsec.StrByRuneNotIn("\n", 0),
			parsec.Rune('\n'),
		),
		func(r parsec.ParseResult) (parsec.ParseResult, error) {
			note := &Note{r.Value().(string)}
			return parsec.NewParseResult(r.SourceRange(), note), nil
		},
	)

	notes := parsec.MapValueAsResults(
		parsec.ManyWithoutNil(
			parsec.Or(
				note,
				parsec.ToNil(blankLine),
			),
		),
		func(rs []parsec.ParseResult) (interface{}, error) {
			notes := make([]*Note, 0, len(rs))
			for _, ret := range rs {
				notes = append(notes, ret.Value().(*Note))
			}
			return notes, nil
		},
	)

	piece := parsec.MapValue(
		parsec.Or(
			parsec.RuneIn("歩香桂銀金飛角王玉龍馬と"),
			parsec.StrByAnd(
				parsec.Rune('成'),
				parsec.RuneIn("香桂銀"),
			),
		),
		func(v interface{}) (interface{}, error) {
			return shogi.MakePieceFromJa(v.(string)), nil
		},
	)

	move := parsec.MapValueAsResults(
		parsec.And(
			parsec.RuneIn("▲△"),
			parsec.OrWithIndex(
				parsec.StrByAnd(
					parsec.RuneIn("１２３４５６７８９"),
					parsec.RuneIn("一二三四五六七八九"),
				),
				parsec.And(
					parsec.Rune('同'),
					parsec.Maybe(parsec.Rune('　')),
				),
			),
			piece,
			parsec.ManyMinMax(parsec.RuneIn("打上引寄右左直"), 0, 2), // not strictly
			parsec.Or(
				parsec.Rune('成'),
				parsec.Str("不成"),
				parsec.Null(),
			),
		),
		func(rs []parsec.ParseResult) (interface{}, error) {
			m := &Move{}
			if rs[0].Value().(string) == "▲" {
				m.Side = shogi.Black
			} else {
				m.Side = shogi.White
			}

			orv := rs[1].Value().(parsec.OrResultValue)
			if orv.Index() == 0 {
				cs := []rune(orv.Value().(string))
				m.Destination = shogi.Position{X: mbnum[cs[0]], Y: mbnum[cs[1]]}
			} else {
				m.Same = true
			}

			m.Piece = rs[2].Value().(shogi.Piece)

			mvdescs := make([]MovementDescriptor, 0, 0)
			for _, ret := range rs[3].Value().([]parsec.ParseResult) {
				c := []rune(ret.Value().(string))[0]
				mvdescs = append(mvdescs, rune2mvdesc[c])
			}
			m.MovementDescriptors = mvdescs

			m.Promotion = false
			if s, ok := rs[4].Value().(string); ok {
				if s == "成" {
					m.Promotion = true
				}
			}

			return m, nil
		},
	)

	moves := parsec.MapValueAsResults(
		parsec.ManyWithoutNil(
			parsec.Or(
				move,
				parsec.ToNil(parsec.RuneIn(" \n")),
			),
		),
		func(rs []parsec.ParseResult) (interface{}, error) {
			moves := make([]*Move, 0, len(rs))
			for _, ret := range rs {
				moves = append(moves, ret.Value().(*Move))
			}
			return moves, nil
		},
	)

	result := parsec.MapValue(
		parsec.StrByAnd(
			parsec.Str("まで"),
			parsec.StrByRuneNotIn("\n", 1),
			parsec.Rune('\n'),
		),
		func(v interface{}) (interface{}, error) {
			return &Result{v.(string)}, nil
		},
	)

	ki2 := parsec.MapValueAsResults(
		parsec.And(
			metaEntries,
			notes,
			moves,
			result,
		),
		func(rs []parsec.ParseResult) (interface{}, error) {
			k := &Ki2{
				MetaEntries: rs[0].Value().([]*MetaEntry),
				Notes:       rs[1].Value().([]*Note),
				Moves:       rs[2].Value().([]*Move),
				Result:      rs[3].Value().(*Result),
			}
			return k, nil
		},
	)

	ret, err := ki2.Parse(r)
	if err != nil {
		return nil, err
	}
	return ret.Value().(*Ki2), nil
}
