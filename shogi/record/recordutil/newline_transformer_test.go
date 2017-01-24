package recordutil

import (
	"bufio"
	"io/ioutil"
	"strings"
	"testing"

	"golang.org/x/text/transform"
)

func TestNewlineNormalizer(t *testing.T) {
	type Case struct {
		Input    string
		Expected string
	}
	cases := []Case{
		Case{
			Input:    "foo\rfoo\r\nfoo\nfoo",
			Expected: "foo\nfoo\nfoo\nfoo",
		},
	}
	for _, c := range cases {
		reader := bufio.NewReader(transform.NewReader(strings.NewReader(c.Input), NewNewlineNormalizer()))
		data, _ := ioutil.ReadAll(reader)
		actual := string(data)
		if actual != c.Expected {
			t.Fatalf("expected: %s, actual: %s", c.Expected, actual)
		}
	}
}
