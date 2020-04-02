package plugfind

import (
	"path"
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
		plug("x/y/z"),
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

	exp = "x/y/z"
	p = f.Find(path.Base(exp), plugs)
	if p == nil {
		t.Fatalf("Expected to find plug %s", exp)
	}

	act = p.PluginName()
	if act != exp {
		t.Fatalf("Expected %s, got %s", exp, act)
	}
}
