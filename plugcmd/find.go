package plugcmd

import (
	"path"
	"strings"

	"github.com/gobuffalo/plugins"
	"github.com/gobuffalo/plugins/plugfind"
)

// FindFromArgs uses the first arg that does not begin with `-`
// as the name argument for Find
func FindFromArgs(args []string, plugs []plugins.Plugin) Commander {
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			continue
		}
		return Find(a, plugs)
	}
	return nil
}

// Find wraps the other cmd finders into a mega finder for cmds
func Find(name string, plugs []plugins.Plugin) Commander {
	for _, p := range plugs {
		c, ok := p.(Commander)
		if !ok {
			continue
		}
		if n, ok := c.(Namer); ok {
			if n.CmdName() == name {
				return c
			}
		}

		if n, ok := c.(Aliaser); ok {
			for _, a := range n.CmdAliases() {
				if a == name {
					return c
				}
			}
		}

		if name == path.Base(c.PluginName()) {
			return c
		}
	}
	return nil
}

// DEPRECRATED: ByAliaser can be used to search plugins that implement
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

// DEPRECRATED: ByNamer can be used to search for plugins that implement
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

// DEPRECRATED: ByCommander can be used to search for plugins that implement
// Commander and who's plugin name matches the name
func ByCommander(f plugfind.Finder) plugfind.FinderFn {
	fn := func(name string, plugs []plugins.Plugin) plugins.Plugin {
		p := f.Find(name, plugs)
		if p == nil {
			return nil
		}
		if _, ok := p.(Commander); ok {
			return p
		}
		return nil
	}
	return plugfind.FinderFn(fn)
}
