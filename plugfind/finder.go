package plugfind

import (
	"github.com/gobuffalo/plugins"
)

type Finder interface {
	Find(name string, plugs []plugins.Plugin) plugins.Plugin
}

type FinderFn func(name string, plugs []plugins.Plugin) plugins.Plugin

func (f FinderFn) Find(name string, plugs []plugins.Plugin) plugins.Plugin {
	return f(name, plugs)
}

func Background() Finder {
	fn := func(name string, plugs []plugins.Plugin) plugins.Plugin {
		for _, p := range plugs {
			if name == p.PluginName() {
				return p
			}
		}
		return nil
	}
	return FinderFn(fn)
}
