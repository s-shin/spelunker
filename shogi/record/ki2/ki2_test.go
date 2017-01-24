package ki2

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFoo(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(wd, "fixtures/2chkifu/40000/53500.ki2")
	_, _ = Load(path)
}
