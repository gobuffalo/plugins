package plugflag

import (
	"flag"
	"fmt"
	"path"

	"github.com/gobuffalo/plugins"
)

// CleanFlagSet sanitizes, and namespaces, flags to be used
// when incorporating flags from other plugins.
func CleanSet(p plugins.Plugin, set *flag.FlagSet) []*flag.Flag {
	var flags []*flag.Flag
	set.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	return Clean(p, flags)
}

// CleanFlags sanitizes, and namespaces, flags to be used
// when incorporating flags from other plugins.
func Clean(p plugins.Plugin, flags []*flag.Flag) []*flag.Flag {
	fls := make([]*flag.Flag, len(flags))
	for i, f := range flags {
		fls[i] = &flag.Flag{
			DefValue: f.DefValue,
			Name:     fmt.Sprintf("%s-%s", path.Base(name(p)), f.Name),
			Usage:    f.Usage,
			Value:    f.Value,
		}
	}
	return fls
}

func name(p plugins.Plugin) string {
	type cmdName interface {
		CmdName() string
	}
	if c, ok := p.(cmdName); ok {
		return c.CmdName()
	}
	return p.PluginName()
}
