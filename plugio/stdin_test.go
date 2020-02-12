package plugio

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/gobuffalo/plugins"
)

func Test_Stdin(t *testing.T) {
	exp := "hello"
	in := strings.NewReader(exp)

	stdin := Stdin([]plugins.Plugin{
		NewInner(in),
	})

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
	act := Stdin(nil)
	if act != exp {
		t.Fatalf("Expected %v, got %v", exp, act)
	}
}
