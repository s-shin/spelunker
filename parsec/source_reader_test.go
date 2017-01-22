package parsec

import (
	"reflect"
	"testing"

	"github.com/k0kubun/pp"
)

func TestSourceReaderCurrentPosition(t *testing.T) {
	r := NewInMemorySourceReader("ab\nc\n\nd\n")

	//    1  2  3
	// 1: a  b  \n
	// 2: c  \n
	// 3: \n
	// 4: d  \n

	type Expected struct {
		Pos         SourcePosition
		IsPosError  bool
		Char        rune
		IsReadError bool
	}

	for _, e := range []Expected{
		{SourcePosition{1, 1}, false, 'a', false},
		{SourcePosition{1, 2}, false, 'b', false},
		{SourcePosition{1, 3}, false, '\n', false},
		{SourcePosition{2, 1}, false, 'c', false},
		{SourcePosition{2, 2}, false, '\n', false},
		{SourcePosition{3, 1}, false, '\n', false},
		{SourcePosition{4, 1}, false, 'd', false},
		{SourcePosition{4, 2}, false, '\n', false},
		{SourcePositionNull, true, 0, true},
	} {
		if p, err := r.CurrentPosition(); !reflect.DeepEqual(e.Pos, p) || e.IsPosError != (err != nil) {
			t.Fatal(pp.Sprintf("expected: %v\nactual: %v, %v", e, p, err))
		}
		if c, err := r.Read(); c != e.Char || e.IsReadError != (err != nil) {
			t.Fatal(pp.Sprintf("expected: %v\nactual: %v, %v", e, c, err))
		}
	}
}

func TestSourceReaderTransaction(t *testing.T) {
	r := NewInMemorySourceReader("abcdefg")

	r.Begin()
	r.Read()
	r.Rollback()
	if p, err := r.CurrentPosition(); p.Column != 1 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	r.Begin()
	r.Read()
	r.Commit()
	if p, err := r.CurrentPosition(); p.Column != 2 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	func() {
		t := r.Transaction()
		defer t.Guard()
		r.Read()
	}()
	if p, err := r.CurrentPosition(); p.Column != 2 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	func() {
		t := r.Transaction()
		defer t.Guard()
		r.Read()
		t.Commit()
	}()
	if p, err := r.CurrentPosition(); p.Column != 3 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	r.Begin()
	r.Begin()
	r.Read()
	r.Commit()
	if p, err := r.CurrentPosition(); p.Column != 4 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}
	r.Rollback()
	if p, err := r.CurrentPosition(); p.Column != 3 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	func() {
		txn := r.Transaction()
		defer txn.Guard()
		r.Read()
		if p, err := r.CurrentPosition(); p.Column != 4 || err != nil {
			t.Fatal(pp.Sprintf("%v, %v", p, err))
		}
		func() {
			txn2 := r.Transaction()
			defer txn2.Guard()
			r.Read()
			if p, err := r.CurrentPosition(); p.Column != 5 || err != nil {
				t.Fatal(pp.Sprintf("%v, %v", p, err))
			}
		}()
		if p, err := r.CurrentPosition(); p.Column != 4 || err != nil {
			t.Fatal(pp.Sprintf("%v, %v", p, err))
		}
		txn.Commit()
	}()
	if p, err := r.CurrentPosition(); p.Column != 4 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	func() {
		txn := r.Transaction()
		r.Read()
		txn.Rollback()
	}()
	if p, err := r.CurrentPosition(); p.Column != 4 || err != nil {
		t.Fatal(pp.Sprintf("%v, %v", p, err))
	}

	if r.Commit() == nil || r.Rollback() == nil {
		t.Fail()
	}
}
