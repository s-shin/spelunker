package parsec

import "fmt"

func ExampleNull() {
	r, err := Null().Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value())
	// Output: <nil>
}

func ExampleAnyRune() {
	r, err := AnyRune().Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: f
}

func ExampleRune() {
	r, err := Rune('f').Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: f
}

func ExampleRuneIn() {
	r, err := RuneIn("abcdef").Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: f
}

func ExampleRuneNotIn() {
	r, err := RuneNotIn("abcde").Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: f
}

func ExampleAnd() {
	r, err := And(Rune('f'), Rune('o')).Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().([]ParseResult)[0].Value().(string))
	// Output: f
}

func ExampleOr() {
	r, err := Or(Rune('a'), Rune('f')).Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: f
}

func ExampleMany() {
	{
		r, err := Many(Rune('a')).Parse(NewInMemorySourceReader("bcde"))
		if err != nil {
			panic(err)
		}
		fmt.Println(len(r.Value().([]ParseResult)))
	}
	{
		r, err := Many(Rune('a')).Parse(NewInMemorySourceReader("aaaaabcde"))
		if err != nil {
			panic(err)
		}
		fmt.Println(len(r.Value().([]ParseResult)))
	}
	// Output:
	// 0
	// 5
}

func ExampleMap() {
	r, err := Map(
		Null(),
		func(r ParseResult) (ParseResult, error) {
			return NewParseResult(r.SourceRange(), "foo"), nil
		},
	).Parse(NewInMemorySourceReader(""))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: foo
}

func ExampleLazy() {
	var foo, foobar Parser
	foobar = Lazy(func(this Parser) Parser {
		// this block will be executed in first call of foobar.Parse.
		return ConcatStr(And(foo, Rune('b'), Rune('a'), Rune('r')))
	})
	foo = Lazy(func(this Parser) Parser {
		return ConcatStr(And(Rune('f'), Rune('o'), Rune('o')))
	})
	r, err := foobar.Parse(NewInMemorySourceReader("foobar"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: foobar
}
