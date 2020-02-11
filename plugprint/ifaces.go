package plugprint

import (
	"io"

	"github.com/gobuffalo/plugins"
	"github.com/gobuffalo/plugins/plugcmd"
)

type Aliaser = plugcmd.Aliaser
type SubCommander = plugcmd.SubCommander
type CommandNamer = plugcmd.Namer
type PluginScoper = plugins.Scoper

// Describer is called by `Print` and can be
// implemented to print a short, single line,
// description of the plugin. `-h`
type Describer interface {
	Description() string
}

type FlagPrinter interface {
	PrintFlags(w io.Writer) error
}

type Hider interface {
	HidePlugin()
}

// UsagePrinter is called by `Print` and can be implemented
// to print a large block of usage information after the
// `Describer` interface is called. This is useful for printing
// flag information, links, and other messages to users.
type UsagePrinter interface {
	PrintUsage(w io.Writer) error
}
