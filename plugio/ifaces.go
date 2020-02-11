package plugio

import (
	"io"

	"github.com/gobuffalo/plugins"
)

type Inner interface {
	plugins.Plugin
	Stdin() io.Reader
}

type Outer interface {
	plugins.Plugin
	Stdout() io.Writer
}

type Errer interface {
	plugins.Plugin
	Stderr() io.Writer
}

type InNeeder interface {
	SetStdin(io.Reader) error
}

type OutNeeder interface {
	SetStdout(io.Writer) error
}

type ErrNeeder interface {
	SetStderr(io.Writer) error
}
