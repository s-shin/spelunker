package parsec

import (
	"container/list"
	"fmt"

	"github.com/pkg/errors"
)

var ErrEOF = errors.New("EOF")

var SourcePositionNull = SourcePosition{}

// SourcePosition shows a position in a source code.
type SourcePosition struct {
	Line   int
	Column int
}

func (p SourcePosition) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

func (p SourcePosition) To(pos SourcePosition) SourceRange {
	return SourceRange{p, pos}
}

var SourceRangeNull = SourceRange{}

// SourceRange shows a range in a source code.
type SourceRange struct {
	Start SourcePosition
	End   SourcePosition
}

func (r SourceRange) String() string {
	return fmt.Sprintf("[%s,%s]", r.Start, r.End)
}

type Transactional interface {
	Begin()
	Commit() error
	Rollback() error
}

// Transaction is a helper of a Transactional object.
type Transaction struct {
	t    Transactional
	done bool
}

// NewTransaction creates a Transaction instance.
func NewTransaction(t Transactional) *Transaction {
	new := &Transaction{t, false}
	new.t.Begin()
	return new
}

func (t *Transaction) Guard() {
	if !t.done {
		_ = t.t.Rollback()
	}
}

func (t *Transaction) Commit() error {
	t.done = true
	return t.t.Commit()
}

func (t *Transaction) Rollback() error {
	t.done = true
	return t.t.Rollback()
}

type SourceReader interface {
	CurrentPosition() (SourcePosition, error)
	Read() (rune, error)
	Transaction() *Transaction
	Begin()
	Commit() error
	Rollback() error
}

type InMemorySourceReader struct {
	runes  []rune
	i      int
	iStack *list.List
}

// NewInMemorySourceReader create an InMemorySourceReader instance.
// Each newline of the source is preferred to be normalized to single LF.
func NewInMemorySourceReader(source string) SourceReader {
	return &InMemorySourceReader{
		runes:  []rune(source),
		i:      0,
		iStack: list.New(),
	}
}

func (r *InMemorySourceReader) CurrentPosition() (SourcePosition, error) {
	if r.i == len(r.runes) {
		return SourcePositionNull, ErrEOF
	}
	line := 1
	lastBeginingOfLine := 0
	var prev rune
	for i, c := range r.runes[:r.i+1] {
		if prev == '\n' {
			line++
			lastBeginingOfLine = i
		}
		prev = c
	}
	return SourcePosition{line, r.i - lastBeginingOfLine + 1}, nil
}

func (r *InMemorySourceReader) Read() (rune, error) {
	if r.i == len(r.runes) {
		return 0, ErrEOF
	}
	c := r.runes[r.i]
	r.i++
	return c, nil
}

// Transaction creates a transaction.
func (r *InMemorySourceReader) Transaction() *Transaction {
	return NewTransaction(r)
}

// Begin is a part of Transactional implementation.
func (r *InMemorySourceReader) Begin() {
	r.iStack.PushBack(r.i)
}

// Commit is a part of Transactional implementation.
func (r *InMemorySourceReader) Commit() error {
	if r.iStack.Len() == 0 {
		return errors.New("no active transaction")
	}
	_ = r.iStack.Remove(r.iStack.Back())
	return nil
}

// Rollback is a part of Transactional implementation.
func (r *InMemorySourceReader) Rollback() error {
	if r.iStack.Len() == 0 {
		return errors.New("no active transaction")
	}
	v := r.iStack.Remove(r.iStack.Back())
	r.i = v.(int)
	return nil
}
