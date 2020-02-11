package plugfind

import (
	"testing"

	"github.com/gobuffalo/plugins"
)

type plug string

func (p plug) PluginName() string {
	return string(p)
}

func Test_Finder(t *testing.T) {
	f := Background()

	plugs := []plugins.Plugin{
		plug("a"),
		plug("b"),
		plug("c"),
	}

	exp := "b"
	p := f.Find(exp, plugs)
	if p == nil {
		t.Fatalf("Expected to find plug %s", exp)
	}

	act := p.PluginName()
	if act != exp {
		t.Fatalf("Expected %s, got %s", exp, act)
	}
}
