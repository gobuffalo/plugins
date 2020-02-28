package plugins

import (
	"sort"
	"strings"
	"testing"
)

func Test_Plugins_Sort(t *testing.T) {
	plugs := Plugins{
		back("b"),
		back("c"),
		back("a"),
	}

	sort.Sort(plugs)

	exp := []string{"a", "b", "c"}
	var act []string
	for _, p := range plugs {
		act = append(act, p.PluginName())
	}

	if strings.Join(act, ",") != strings.Join(exp, ",") {
		t.Fatalf("Expected %q, got %q", exp, act)
	}
}

func Test_Sort(t *testing.T) {
	plugs := []Plugin{
		back("b"),
		back("c"),
		back("a"),
	}

	Sort(plugs)

	exp := []string{"a", "b", "c"}
	var act []string
	for _, p := range plugs {
		act = append(act, p.PluginName())
	}

	if strings.Join(act, ",") != strings.Join(exp, ",") {
		t.Fatalf("Expected %q, got %q", exp, act)
	}
}
