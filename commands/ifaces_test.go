package commands

import (
	"context"

	"github.com/gobuffalo/plugins"
)

type background string

func (b background) PluginName() string {
	return string(b)
}

var _ plugins.Plugin = aliaser{}
var _ Aliaser = aliaser{}

type aliaser []string

func (a aliaser) PluginName() string {
	return "aliaser"
}

func (a aliaser) CmdAliases() []string {
	return []string(a)
}

var _ plugins.Plugin = namer("")
var _ Namer = namer("")

type namer string

func (namer) PluginName() string {
	return "namer"
}

func (n namer) CmdName() string {
	return string(n)
}

var _ plugins.Plugin = commander(nil)
var _ Commander = commander(nil)

func (c commander) PluginName() string {
	return "commander"
}

type commander func(ctx context.Context, root string, args []string) error

func (c commander) Main(ctx context.Context, root string, args []string) error {
	return c(ctx, root, args)
}
