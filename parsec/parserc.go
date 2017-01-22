package parsec

import "fmt"

type ParseError interface {
	Error() string
	SourceRange() SourceRange
}

type parseError struct {
	r   SourceRange
	msg string
}

func NewParseError(r SourceRange, msg string) ParseError {
	return &parseError{r, msg}
}

func NewParseErrorExpect(r SourceRange, expected, found string) ParseError {
	return NewParseError(r, fmt.Sprintf("expected %s, found %s %s", expected, found, r))
}

func (e *parseError) SourceRange() SourceRange {
	return e.r
}

func (e *parseError) Error() string {
	return e.msg
}

//---

type ParseResult interface {
	SourceRange() SourceRange
	Value() interface{}
	IsNull() bool
}

type parseResult struct {
	r SourceRange
	v interface{}
}

var ParseResultNull = NewParseResult(SourceRangeNull, nil)

// NewParseResult creates a common ParseResult.
func NewParseResult(r SourceRange, v interface{}) ParseResult {
	return &parseResult{r, v}
}

func NewParseResultFromResults(rs []ParseResult) ParseResult {
	var start, end SourcePosition
	for i := 0; i < len(rs); i++ {
		pos := rs[i].SourceRange().Start
		if pos != SourcePositionNull {
			start = pos
			break
		}
	}
	for i := len(rs) - 1; i >= 0; i-- {
		pos := rs[i].SourceRange().End
		if pos != SourcePositionNull {
			end = pos
			break
		}
	}
	return NewParseResult(SourceRange{start, end}, rs)
}

func (r *parseResult) SourceRange() SourceRange {
	return r.r
}

func (r *parseResult) Value() interface{} {
	return r.v
}

func (r *parseResult) IsNull() bool {
	return r.r == SourceRangeNull
}

//---

type Parser interface {
	Parse(r SourceReader) (ParseResult, error)
	// String should be return what parser do.
	String() string
}
