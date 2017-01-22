package parsec

import (
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"
)

//------------------------------------------------------------------------------
// Utilities
//------------------------------------------------------------------------------

func runemap(rs []rune) map[rune]struct{} {
	m := make(map[rune]struct{})
	for _, r := range rs {
		m[r] = struct{}{}
	}
	return m
}

//------------------------------------------------------------------------------
// Primitive Parsers
//------------------------------------------------------------------------------

type nullParser struct{}

func (p *nullParser) String() string {
	return "Null"
}

func (p *nullParser) Parse(r SourceReader) (ParseResult, error) {
	return ParseResultNull, nil
}

func Null() Parser {
	return &nullParser{}
}

//---

type anyRuneParser struct{}

// AnyRune returns a parser to parse any character.
func AnyRune() Parser {
	return &anyRuneParser{}
}

func (p *anyRuneParser) String() string {
	return "AnyRune"
}

func (p *anyRuneParser) Parse(r SourceReader) (ParseResult, error) {
	txn := r.Transaction()
	defer txn.Guard()
	pos, err := r.CurrentPosition()
	if err != nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
	}
	c, err := r.Read()
	if err != nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
	}
	txn.Commit()
	return NewParseResult(SourceRange{pos, pos}, string(c)), nil
}

//---

type runeInParser struct {
	accept map[rune]struct{}
	not    bool
}

func (p *runeInParser) String() string {
	var s string
	for r := range p.accept {
		s += string(r)
	}
	var not string
	if p.not {
		not = "Not"
	}
	return "Rune" + not + "In(" + s + ")"
}

func (p *runeInParser) Parse(r SourceReader) (ParseResult, error) {
	txn := r.Transaction()
	defer txn.Guard()
	pos, err := r.CurrentPosition()
	if err != nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
	}
	c, err := r.Read()
	if err != nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
	}
	if _, ok := p.accept[c]; ok == p.not {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("'%s'", string(c)))
	}
	txn.Commit()
	return NewParseResult(SourceRange{pos, pos}, string(c)), nil
}

func RuneIn(s string) Parser {
	return &runeInParser{runemap([]rune(s)), false}
}

func RuneNotIn(s string) Parser {
	return &runeInParser{runemap([]rune(s)), true}
}

func Rune(c rune) Parser {
	return &runeInParser{runemap([]rune{c}), false}
}

//---

type andParser struct {
	children []Parser
}

func (p *andParser) String() string {
	var s string
	for _, child := range p.children {
		if len(s) > 0 {
			s += ", "
		}
		s += child.String()
	}
	return "And(" + s + ")"
}

func (p *andParser) Parse(r SourceReader) (ParseResult, error) {
	txn := r.Transaction()
	defer txn.Guard()
	pos, err := r.CurrentPosition()
	if err != nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
	}
	rs := make([]ParseResult, 0, len(p.children))
	for _, child := range p.children {
		cur, _ := r.CurrentPosition()
		ret, err := child.Parse(r)
		if err != nil {
			return nil, NewParseErrorExpect(SourceRange{pos, cur}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
		}
		rs = append(rs, ret)
	}
	txn.Commit()
	return NewParseResultFromResults(rs), nil
}

// And creates sequencial parser.
func And(ps ...Parser) Parser {
	return &andParser{ps}
}

//---

type orParser struct {
	children []Parser
}

func (p *orParser) String() string {
	var s string
	for _, child := range p.children {
		if len(s) > 0 {
			s += ", "
		}
		s += child.String()
	}
	return "Or(" + s + ")"
}

func (p *orParser) Parse(r SourceReader) (ParseResult, error) {
	txn := r.Transaction()
	defer txn.Guard()
	pos, err := r.CurrentPosition()
	if err != nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), fmt.Sprintf("error (%s)", err.Error()))
	}
	var result ParseResult
	for _, child := range p.children {
		ret, err := child.Parse(r)
		if err != nil {
			continue
		}
		result = ret
		break
	}
	if result == nil {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), "none to be parsed")
	}
	txn.Commit()
	return result, nil
}

func Or(ps ...Parser) Parser {
	return &orParser{ps}
}

//---

type manyParser struct {
	child Parser
	least int
}

func (p *manyParser) String() string {
	return fmt.Sprintf("Many%d(%s)", p.least, p.child.String())
}

func (p *manyParser) Parse(r SourceReader) (ParseResult, error) {
	txn := r.Transaction()
	defer txn.Guard()
	pos, err := r.CurrentPosition()
	if err != nil {
		return ParseResultNull, nil
	}
	rs := make([]ParseResult, 0, 0)
	for {
		ret, err := p.child.Parse(r)
		if err != nil {
			break
		}
		rs = append(rs, ret)
	}
	if len(rs) < p.least {
		return nil, NewParseErrorExpect(SourceRange{pos, pos}, p.String(), "no sufficient data to parse")
	}
	txn.Commit()
	return NewParseResultFromResults(rs), nil
}

func Many(p Parser) Parser {
	return &manyParser{p, 0}
}

func Many1(p Parser) Parser {
	return &manyParser{p, 1}
}

func ManyN(p Parser, n int) Parser {
	return &manyParser{p, n}
}

//---

type MapFunc func(r ParseResult, err error) (ParseResult, error)

type mapParser struct {
	child Parser
	fn    MapFunc
}

func (p *mapParser) String() string {
	return "Map"
}

func (p *mapParser) Parse(r SourceReader) (ParseResult, error) {
	ret, err := p.child.Parse(r)
	return p.fn(ret, err)
}

func MapWithError(p Parser, fn MapFunc) Parser {
	return &mapParser{p, fn}
}

//---

type lazyParser struct {
	child Parser
	init  func(this Parser) Parser
}

func (p *lazyParser) String() string {
	return "Lazy"
}

func (p *lazyParser) Parse(r SourceReader) (ParseResult, error) {
	if p.child == nil {
		p.child = p.init(p)
	}
	return p.child.Parse(r)
}

func Lazy(init func(this Parser) Parser) Parser {
	return &lazyParser{nil, init}
}

//------------------------------------------------------------------------------
// High-level Combinators
//------------------------------------------------------------------------------

func Map(p Parser, fn func(r ParseResult) (ParseResult, error)) Parser {
	return MapWithError(p, func(r ParseResult, err error) (ParseResult, error) {
		if err == nil {
			return fn(r)
		}
		return r, err
	})
}

func Filter(p Parser, filter func(r ParseResult) bool) Parser {
	return Map(p, func(r ParseResult) (ParseResult, error) {
		if results, ok := r.Value().([]ParseResult); ok {
			filtered := make([]ParseResult, 0, len(results))
			for _, result := range results {
				if filter(result) {
					filtered = append(filtered, result)
				}
			}
			return NewParseResult(r.SourceRange(), filtered), nil
		}
		return nil, NewParseError(r.SourceRange(), "expected []ParseResult")
	})
}

// FilterNullResult filters null results (e.g. ParseResultNull).
func FilterNullResult(p Parser) Parser {
	return Filter(p, func(r ParseResult) bool {
		return !r.IsNull()
	})
}

// FilterNil filters results the value of which is nil.
// It also removes null results.
func FilterNil(p Parser) Parser {
	return Filter(FilterNullResult(p), func(r ParseResult) bool {
		return r.Value() != nil
	})
}

// ConcatStr converts the result values to one concatenated string.
func ConcatStr(p Parser) Parser {
	return Map(p, func(r ParseResult) (ParseResult, error) {
		if results, ok := r.Value().([]ParseResult); ok {
			var str string
			for _, result := range results {
				if s, ok := result.Value().(string); ok {
					str += s
				} else {
					return nil, NewParseError(result.SourceRange(), "not string")
				}
			}
			return NewParseResult(r.SourceRange(), str), nil
		}
		if str, ok := r.Value().(string); ok {
			return NewParseResult(r.SourceRange(), str), nil
		}
		return nil, NewParseError(r.SourceRange(), "failed to concat strings")
	})
}

func Digit() Parser {
	return RuneIn("0123456789")
}

func Digits() Parser {
	return ConcatStr(Many1(Digit()))
}

// ToInt converts the result value to int.
func ToInt(p Parser) Parser {
	return Map(p, func(r ParseResult) (ParseResult, error) {
		s, ok := r.Value().(string)
		if !ok {
			return nil, NewParseError(r.SourceRange(), "failed to convert value to int")
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, NewParseError(r.SourceRange(), "failed to convert value to int")
		}
		return NewParseResult(r.SourceRange(), i), nil
	})
}

// ToNil sets nil as the result value.
func ToNil(p Parser) Parser {
	return Map(p, func(r ParseResult) (ParseResult, error) {
		return NewParseResult(r.SourceRange(), nil), nil
	})
}

func Index(p Parser, n int) Parser {
	return Map(p, func(r ParseResult) (ParseResult, error) {
		rs, ok := r.Value().([]ParseResult)
		if !ok {
			return nil, NewParseError(r.SourceRange(), "value is not an array")
		}
		if n >= len(rs) {
			return nil, NewParseError(r.SourceRange(), "out of range")
		}
		return NewParseResult(r.SourceRange(), rs[n].Value()), nil
	})
}

// Dump prints the parse result and error for debug.
func Dump(p Parser) Parser {
	return MapWithError(p, func(r ParseResult, err error) (ParseResult, error) {
		pp.Println(r, err)
		return r, err
	})
}
