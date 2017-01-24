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

func ExampleOrWithIndex() {
	r, err := OrWithIndex(Rune('a'), Rune('f')).Parse(NewInMemorySourceReader("foo"))
	if err != nil {
		panic(err)
	}
	v := r.Value().(OrResultValue)
	fmt.Println(v.Index(), v.Value().(string))
	// Output: 1 f
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

func ExampleMaybe() {
	{
		r, err := Maybe(Rune('f')).Parse(NewInMemorySourceReader("foo"))
		if err != nil {
			panic(err)
		}
		fmt.Println(r.IsNull(), r.Value().(string))
	}
	{
		r, err := Maybe(Rune('o')).Parse(NewInMemorySourceReader("foo"))
		if err != nil {
			panic(err)
		}
		fmt.Println(r.IsNull(), r.Value())
	}
	// Output:
	// false f
	// true <nil>
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

func ExampleStr() {
	r, err := Str("foo").Parse(NewInMemorySourceReader("foobar"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: foo
}

func ExampleStrByAnd() {
	r, err := StrByAnd(Rune('f'), Rune('o')).Parse(NewInMemorySourceReader("foobar"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: fo
}

func ExampleStrByRuneIn() {
	r, err := StrByRuneIn("fo", 2).Parse(NewInMemorySourceReader("foobar"))
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Value().(string))
	// Output: foo
}
