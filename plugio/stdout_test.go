package plugio

import (
	"bytes"
	"os"
	"testing"
)

func Test_Stdout(t *testing.T) {
	bb := &bytes.Buffer{}

	stdout := Stdout(NewOuter(bb))

	exp := "hi"
	stdout.Write([]byte(exp))

	act := bb.String()

	if act != exp {
		t.Fatalf("Expected %s, got %s", exp, act)
	}
}

func Test_Stdout_default(t *testing.T) {
	exp := os.Stdout
	act := Stdout()
	if act != exp {
		t.Fatalf("Expected %v, got %v", exp, act)
	}
}
