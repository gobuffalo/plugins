package plugcmd

import (
	"context"

	"github.com/gobuffalo/plugins"
)

// Commander is a plugin that is meant to be the beginning of
// a CLI application
type Commander interface {
	plugins.Plugin
	Main(ctx context.Context, root string, args []string) error
}

// Aliaser is a command that provides aliases for itself
type Aliaser interface {
	plugins.Plugin
	CmdAliases() []string
}

// Namer is a command that provides a different name for the
// command than the name of the Plugin
type Namer interface {
	plugins.Plugin
	CmdName() string
}

// SubCommander can be implemented to provide a list of plugins.Plugin
// that can be used as sub-commands of the current Plugin
type SubCommander interface {
	SubCommands() []plugins.Plugin
}
