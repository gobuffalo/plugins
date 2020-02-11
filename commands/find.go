package commands

import (
	"github.com/gobuffalo/plugins"
	"github.com/gobuffalo/plugins/plugfind"
)

// ByAliaser can be used to search plugins that implement
// Aliaser and who's alias matches the name
func ByAliaser(f plugfind.Finder) plugfind.Finder {
	fn := func(name string, plugs []plugins.Plugin) plugins.Plugin {
		for _, p := range plugs {
			if a, ok := p.(Aliaser); ok {
				for _, ax := range a.CmdAliases() {
					if ax == name {
						return p
					}
				}
			}
		}
		return f.Find(name, plugs)
	}
	return plugfind.FinderFn(fn)
}

// ByNamer can be used to search for plugins that implement
// Namer and who's CmdName matches the name
func ByNamer(f plugfind.Finder) plugfind.Finder {
	fn := func(name string, plugs []plugins.Plugin) plugins.Plugin {
		for _, p := range plugs {
			if c, ok := p.(Namer); ok {
				if c.CmdName() == name {
					return p
				}
			}
		}
		return f.Find(name, plugs)
	}
	return plugfind.FinderFn(fn)
}

// ByCommander can be used to search for plugins that implement
// Commander and who's plugin name matches the name
func ByCommander(f plugfind.Finder) plugfind.FinderFn {
	fn := func(name string, plugs []plugins.Plugin) plugins.Plugin {
		p := f.Find(name, plugs)
		if p == nil {
			return nil
		}
		if c, ok := p.(Commander); ok {
			if c.PluginName() == name {
				return p
			}
		}
		return nil
	}
	return plugfind.FinderFn(fn)
}
