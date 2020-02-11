package plugflag

import (
	"flag"
	"testing"
)

type background string

func (b background) PluginName() string {
	return string(b)
}

func Test_CleanFlags(t *testing.T) {
	p := background("foo")

	flags := Clean(p, []*flag.Flag{
		{Name: "my-flag"},
	})

	if len(flags) != 1 {
		t.Fatalf("Expected %d flags, got %d", 1, len(flags))
	}

	f := flags[0]

	exp := "foo-my-flag"
	act := f.Name
	if act != exp {
		t.Fatalf("Expected %s, got %s)", exp, act)
	}
}

func Test_CleanFlagSet(t *testing.T) {
	p := background("foo")

	set := flag.NewFlagSet("", flag.ContinueOnError)
	set.String("my-flag", "", "")
	flags := CleanSet(p, set)

	if len(flags) != 1 {
		t.Fatalf("Expected %d flags, got %d", 1, len(flags))
	}

	f := flags[0]

	exp := "foo-my-flag"
	act := f.Name
	if act != exp {
		t.Fatalf("Expected %s, got %s)", exp, act)
	}
}
