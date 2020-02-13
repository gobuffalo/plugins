package plugio

import (
	"io"
	"os"

	"github.com/gobuffalo/plugins"
)

var _ Outer = outer{}

type outer struct {
	io.Writer
}

func (i outer) PluginName() string {
	return "outer"
}

func (i outer) Stdout() io.Writer {
	return i.Writer
}

// NewOuter converts an io.Writer to an Outer plugin.
func NewOuter(r io.Writer) Outer {
	return outer{Writer: r}
}

// Stdout returns a io.MultiWriter containing all
// plugins that implement Outer. If none are found,
// then os.Stdout is returned
func Stdout(plugs ...plugins.Plugin) io.Writer {
	var ins []io.Writer

	for _, p := range plugs {
		if x, ok := p.(Outer); ok {
			ins = append(ins, x.Stdout())
		}
	}

	if len(ins) == 0 {
		return os.Stdout
	}
	return io.MultiWriter(ins...)
}
