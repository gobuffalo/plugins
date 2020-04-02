package plugins

import (
	"fmt"
	"testing"
)

func Test_Errors(t *testing.T) {
	t.Parallel()

	err := fmt.Errorf("boom")
	p := back("a")
	err = Wrap(p, err)

	exp := `github.com/gobuffalo/plugins.back (a): boom`
	act := err.Error()
	if act != exp {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}
