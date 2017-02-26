package script

import (
	"reflect"
	"testing"

	"github.com/k0kubun/pp"
)

func TestParseShogiScript(t *testing.T) {
	type Expected struct {
		Commands []*Command
		IsError  bool
	}
	type Case struct {
		Input    string
		Expected Expected
	}
	cases := []Case{
		{
			"foo bar",
			Expected{
				[]*Command{
					{"foo", []*CommandArg{{"", "bar"}}},
				},
				false,
			},
		},
		{
			`
			# comment
			begin aaa:"f\"o\\o"
			end """f"o\"o"""
			`,
			Expected{
				[]*Command{
					{"begin", []*CommandArg{{"aaa", "f\"o\\o"}}},
					{"end", []*CommandArg{{"", "f\"o\"o"}}},
				},
				false,
			},
		},
		{
			`
			foo bbb:"""
			hoge
			"""
			`,
			Expected{
				[]*Command{
					{"foo", []*CommandArg{{"bbb", "\n\t\t\thoge\n\t\t\t"}}},
				},
				false,
			},
		},
	}
	for _, c := range cases {
		cmds, err := ParseShogiScript(c.Input)
		if !reflect.DeepEqual(c.Expected.Commands, cmds) || c.Expected.IsError != (err != nil) {
			t.Fatal(pp.Sprintf("expected: %v\nactual: %v, %v", c.Expected, cmds, err))
		}
	}
}
