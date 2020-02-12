package plugio

import (
	"bytes"
	"os"
	"testing"

	"github.com/gobuffalo/plugins"
)

func Test_Stderr(t *testing.T) {
	bb := &bytes.Buffer{}

	stdout := Stderr([]plugins.Plugin{
		NewErrer(bb),
	})

	exp := "hi"
	stdout.Write([]byte(exp))

	act := bb.String()

	if act != exp {
		t.Fatalf("Expected %s, got %s", exp, act)
	}
}

func Test_Stderr_default(t *testing.T) {
	exp := os.Stderr
	act := Stderr(nil)
	if act != exp {
		t.Fatalf("Expected %v, got %v", exp, act)
	}
}
