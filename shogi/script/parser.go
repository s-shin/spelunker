package script

import "github.com/s-shin/spelunker/parsec"

type CommandArg struct {
	Label, Value string
}

type Command struct {
	Name string
	Args []*CommandArg
}

var shogiScriptParser parsec.Parser

func ShogiScriptParser() parsec.Parser {
	if shogiScriptParser != nil {
		return shogiScriptParser
	}

	ws := parsec.ConstNil(parsec.RuneIn(" \t"))
	ws0 := parsec.ConstNil(parsec.Many(ws))
	ws1 := parsec.ConstNil(parsec.Many1(ws))

	symbol := parsec.StrByAnd(
		parsec.Or(
			parsec.Range('A', 'Z'),
			parsec.Range('a', 'z'),
			parsec.RuneIn("_-"),
		),
		parsec.StrByManyN(
			parsec.Or(
				parsec.Range('A', 'Z'),
				parsec.Range('a', 'z'),
				parsec.Range('0', '9'),
				parsec.RuneIn("_-"),
			),
			1,
		),
	)

	commandName := symbol

	str := parsec.Or(
		// """foo"""
		parsec.StrByAnd(
			parsec.ConstNil(parsec.Str("\"\"\"")),
			parsec.StrByManyNT(
				parsec.Or(
					parsec.ConstStr(parsec.Str("\\\\"), "\\"),
					parsec.ConstStr(parsec.Str("\\\""), "\""),
					parsec.AnyRune(),
				),
				0,
				parsec.ConstNil(parsec.Str("\"\"\"")),
			),
		),
		// "foo"
		parsec.StrByAnd(
			parsec.ConstNil(parsec.Rune('"')),
			parsec.StrByManyNT(
				parsec.Or(
					parsec.ConstStr(parsec.Str("\\\\"), "\\"),
					parsec.ConstStr(parsec.Str("\\\""), "\""),
					parsec.RuneNotIn("\""),
				),
				0,
				parsec.ConstNil(parsec.Rune('"')),
			),
		),
		// foo
		parsec.StrByRuneNotIn(" \t\n", 1),
	)

	commandArg := parsec.MapValueAsResults(
		parsec.And(
			parsec.Maybe(
				parsec.Index(
					parsec.And(
						symbol,
						parsec.Rune(':'),
					),
					0,
				),
			),
			str,
		),
		func(results []parsec.ParseResult) (interface{}, error) {
			var arg CommandArg
			if !results[0].IsNull() {
				arg.Label = results[0].Value().(string)
			}
			arg.Value = results[1].Value().(string)
			return &arg, nil
		},
	)

	commandLine := parsec.MapValueAsResults(
		parsec.And(
			ws0,
			commandName,
			parsec.Many(parsec.Index(parsec.And(ws1, commandArg), 1)),
			ws0,
		),
		func(results []parsec.ParseResult) (interface{}, error) {
			c := &Command{}
			c.Name = results[1].Value().(string)
			argsResults := results[2].Value().([]parsec.ParseResult)
			c.Args = make([]*CommandArg, 0, len(argsResults))
			for _, r := range argsResults {
				c.Args = append(c.Args, r.Value().(*CommandArg))
			}
			return c, nil
		},
	)

	commentLine := parsec.StrByAnd(
		ws0,
		parsec.Rune('#'),
		parsec.StrByRuneNotIn("\n", 0),
	)

	emptyLine := ws0

	line := parsec.Or(
		parsec.ConstNil(commentLine),
		commandLine,
		parsec.ConstNil(emptyLine),
	)

	lines := parsec.ManyWithoutNil(
		parsec.Index(
			parsec.And(
				parsec.Maybe(parsec.Rune('\n')),
				line,
			),
			1,
		),
	)

	shogiScriptParser = lines
	return shogiScriptParser
}

func ParseShogiScript(s string) ([]*Command, error) {
	p := ShogiScriptParser()
	result, err := p.Parse(parsec.NewInMemorySourceReader(s))
	if err != nil {
		return nil, err
	}
	rs := result.Value().([]parsec.ParseResult)
	commands := make([]*Command, 0, len(rs))
	for _, r := range rs {
		commands = append(commands, r.Value().(*Command))
	}
	return commands, nil
}
