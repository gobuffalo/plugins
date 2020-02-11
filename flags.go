package plugins

import (
	"flag"
	"fmt"
	"path"
)

// CleanFlagSet sanitizes, and namespaces, flags to be used
// when incorporating flags from other plugins.
func CleanFlagSet(p Plugin, set *flag.FlagSet) []*flag.Flag {
	var flags []*flag.Flag
	set.VisitAll(func(f *flag.Flag) {
		flags = append(flags, f)
	})
	return CleanFlags(p, flags)
}

// CleanFlags sanitizes, and namespaces, flags to be used
// when incorporating flags from other plugins.
func CleanFlags(p Plugin, flags []*flag.Flag) []*flag.Flag {
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

func name(p Plugin) string {
	type cmdName interface {
		CmdName() string
	}
	if c, ok := p.(cmdName); ok {
		return c.CmdName()
	}
	return p.PluginName()
}
