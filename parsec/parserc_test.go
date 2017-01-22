package parsec

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestParsecSampleCalculator(t *testing.T) {
	var ws, num, expr, expr1, expr2 Parser
	ws = Many(Rune(' '))
	num = ToInt(Digits())

	expr = Lazy(func(this Parser) Parser {
		return Or(
			Map(
				And(ws, expr1, ws, Rune('+'), ws, expr1, ws),
				func(r ParseResult) (ParseResult, error) {
					vs := r.Value().([]ParseResult)
					lv := vs[1].Value().(int)
					rv := vs[5].Value().(int)
					return NewParseResult(r.SourceRange(), lv+rv), nil
				},
			),
			expr1,
		)
	})

	expr1 = Lazy(func(this Parser) Parser {
		return Or(
			Map(
				And(ws, expr2, ws, Rune('*'), ws, expr2, ws),
				func(r ParseResult) (ParseResult, error) {
					vs := r.Value().([]ParseResult)
					lv := vs[1].Value().(int)
					rv := vs[5].Value().(int)
					return NewParseResult(r.SourceRange(), lv*rv), nil
				},
			),
			expr2,
		)
	})

	expr2 = Or(
		Index(And(ws, Rune('('), ws, expr, ws, Rune(')'), ws), 3),
		num,
	)

	calc := expr

	r := NewInMemorySourceReader("10 * 20 + (30 + 40 * 50) * 60")
	result, err := calc.Parse(r)
	if err != nil {
		t.Fatal(pp.Sprint(err))
	}
	if _, err := r.CurrentPosition(); err == nil {
		t.Fatal("not EOF")
	}
	if v, ok := result.Value().(int); !ok || v != 122000 {
		t.Fatal(pp.Sprint(v), pp.Sprint(ok))
	}
}

//---

type Program struct {
	SourceRange
	Comments    []*Comment
	Expressions []Expression
}

type Comment struct {
	SourceRange
	Text string
}

type Expression interface {
	expr()
}

type AssignExpression struct {
	SourceRange
	Identifier *Identifier
	Value      Expression
}

func (e *AssignExpression) expr() {}

type Identifier struct {
	SourceRange
	Name string
}

func (i *Identifier) expr() {}

type Number struct {
	SourceRange
	Value int
}

func (n *Number) expr() {}

type CalculateExpression struct {
	SourceRange
	Left     Expression
	Operator string
	Right    Expression
}

func (e *CalculateExpression) expr() {}

type FunctionCallExpression struct {
	SourceRange
	Function  *Identifier
	Arguments []Expression
}

func (e *FunctionCallExpression) expr() {}

func TestParsecSampleLang(t *testing.T) {
	var wsc, ws, ws1 Parser
	var terminate Parser
	var num, ident, value Parser
	var commentLine Parser
	var calc, calc1, calc2 Parser

	wsc = RuneIn(" \n")
	ws = ToNil(Many(wsc))
	ws1 = ToNil(Many1(wsc))

	terminate = ToNil(Rune(';'))

	num = Map(
		ToInt(Digits()),
		func(r ParseResult) (ParseResult, error) {
			return NewParseResult(r.SourceRange(), &Number{r.SourceRange(), r.Value().(int)}), nil
		})

	ident = Map(
		ConcatStr(Many1(RuneIn("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"))),
		func(r ParseResult) (ParseResult, error) {
			return NewParseResult(r.SourceRange(), &Identifier{r.SourceRange(), r.Value().(string)}), nil
		})

	value = Or(num, ident)

	commentLine = Map(
		ConcatStr(And(Rune('/'), Rune('/'), ConcatStr(Many(RuneNotIn("\n"))), Rune('\n'))),
		func(r ParseResult) (ParseResult, error) {
			return NewParseResult(r.SourceRange(), &Comment{r.SourceRange(), r.Value().(string)}), nil
		})

	calc1 = Lazy(func(this Parser) Parser {
		return Or(
			Map(
				And(calc2, ws, RuneIn("*/"), ws, calc2, ws),
				func(r ParseResult) (ParseResult, error) {
					rs := r.Value().([]ParseResult)
					lv := rs[0].Value().(Expression)
					op := rs[2].Value().(string)
					rv := rs[4].Value().(Expression)
					return NewParseResult(r.SourceRange(), &CalculateExpression{r.SourceRange(), lv, op, rv}), nil
				},
			),
			calc2,
		)
	})

	calc2 = Lazy(func(this Parser) Parser {
		return Or(
			Index(And(Rune('('), ws, calc, ws, Rune(')')), 2),
			value,
		)
	})

	calc = Lazy(func(this Parser) Parser {
		return Or(
			Map(
				And(calc1, ws, RuneIn("+-"), ws, calc1, ws),
				func(r ParseResult) (ParseResult, error) {
					rs := r.Value().([]ParseResult)
					lv := rs[0].Value().(Expression)
					op := rs[2].Value().(string)
					rv := rs[4].Value().(Expression)
					return NewParseResult(r.SourceRange(), &CalculateExpression{r.SourceRange(), lv, op, rv}), nil
				},
			),
			calc1,
		)
	})

	assign := Map(
		And(ident, ws, Rune('='), ws, calc),
		func(r ParseResult) (ParseResult, error) {
			rs := r.Value().([]ParseResult)
			lhs := rs[0].Value().(*Identifier)
			rhs := rs[4].Value().(Expression)
			return NewParseResult(r.SourceRange(), &AssignExpression{r.SourceRange(), lhs, rhs}), nil
		})

	fnargs := Map(
		Many(
			Or(
				Index(And(Rune(','), ws, calc, ws), 2),
				Index(And(calc, ws), 0),
			),
		),
		func(r ParseResult) (ParseResult, error) {
			rs := r.Value().([]ParseResult)
			es := make([]Expression, 0, len(rs))
			for _, ret := range rs {
				es = append(es, ret.Value().(Expression))
			}
			return NewParseResult(r.SourceRange(), es), nil
		})

	fncall := Map(
		And(ident, Rune('('), ws, fnargs, ws, Rune(')')),
		func(r ParseResult) (ParseResult, error) {
			rs := r.Value().([]ParseResult)
			fn := rs[0].Value().(*Identifier)
			args := rs[3].Value().([]Expression)
			return NewParseResult(r.SourceRange(), &FunctionCallExpression{r.SourceRange(), fn, args}), nil
		})

	expr := Or(
		assign,
		fncall,
	)

	program := Map(
		FilterNil(
			Many(
				Or(
					ws1,
					commentLine,
					Index(FilterNil(And(expr, terminate)), 0),
				),
			),
		),
		func(r ParseResult) (ParseResult, error) {
			rs := r.Value().([]ParseResult)
			es := make([]Expression, 0, 0)
			cs := make([]*Comment, 0, 0)
			for _, ret := range rs {
				switch v := ret.Value().(type) {
				case Expression:
					es = append(es, v)
				case *Comment:
					cs = append(cs, v)
				default:
					panic("unknown type")
				}
			}
			return NewParseResult(r.SourceRange(), &Program{r.SourceRange(), cs, es}), nil
		})

	r := NewInMemorySourceReader(`
// comment
foo = 100 + 200;
bar = foo * foo;
fizz_buzz = foo * 2 - bar / (3 - 1);
print(fizz_buzz, 1, 2 + 3);
`)

	// expected := &Program{
	// }

	result, err := program.Parse(r)
	if err != nil {
		t.Fatal(err)
	}
	if ast, ok := result.Value().(*Program); !ok {
		t.Fatal(ok)
	} else {
		pp.Println(ast != nil)
	}
}
