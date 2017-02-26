package script

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestRunnerRunLines(t *testing.T) {
	runner := NewRunner()
	rs, err := runner.RunLines(`
version 1

title "Example 1"

edit ALL -
edit - -51OU
edit - +53KI
edit - +KI

move +0052KI

show state
`)
	if err != nil {
		t.Fatal(pp.Sprint(rs, err))
	}
}
