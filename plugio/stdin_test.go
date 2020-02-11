package plugio

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_Stdin(t *testing.T) {
	exp := "hello"
	in := strings.NewReader(exp)

	stdin := Stdin(NewInner(in))

	b, err := ioutil.ReadAll(stdin)
	if err != nil {
		t.Fatal(err)
	}

	act := string(b)
	if act != exp {
		t.Fatalf("Expected %s, got %s", exp, act)
	}
}

func Test_Stdin_default(t *testing.T) {
	exp := os.Stdin
	act := Stdin()
	if act != exp {
		t.Fatalf("Expected %v, got %v", exp, act)
	}
}
