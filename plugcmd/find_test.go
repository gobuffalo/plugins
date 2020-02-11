package plugcmd

import (
	"testing"

	"github.com/gobuffalo/plugins"
	"github.com/gobuffalo/plugins/plugfind"
)

func Test_ByAliaser(t *testing.T) {
	fn := plugfind.Background()
	fn = ByAliaser(fn)

	plugs := []plugins.Plugin{
		background("a"),
		aliaser{"b"},
		background("c"),
	}

	p := fn.Find("b", plugs)

	if p == nil {
		t.Fatalf("Expected to find plugin")
	}
}

func Test_ByNamer(t *testing.T) {
	fn := plugfind.Background()
	fn = ByNamer(fn)

	plugs := []plugins.Plugin{
		background("a"),
		namer("b"),
		background("c"),
	}

	p := fn.Find("b", plugs)

	if p == nil {
		t.Fatalf("Expected to find plugin")
	}
}

func Test_ByCommander(t *testing.T) {
	plugs := []plugins.Plugin{
		background("a"),
		namer("b"),
		commander(nil),
	}

	fn := plugfind.Background()
	fn = ByCommander(fn)

	p := fn.Find("commander", plugs)
	if p == nil {
		t.Fatalf("Expected to find plugin")
	}

	p = fn.Find("b", plugs)
	if p != nil {
		t.Fatalf("Expected to not find plugin")
	}
}
