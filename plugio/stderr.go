package plugio

import (
	"io"
	"os"

	"github.com/gobuffalo/plugins"
)

var _ Errer = errer{}

type errer struct {
	io.Writer
}

func (i errer) PluginName() string {
	return "errer"
}

func (i errer) Stderr() io.Writer {
	return i.Writer
}

// NewErrer converts an io.Writer to an Outer plugin.
func NewErrer(r io.Writer) Errer {
	return errer{Writer: r}
}

// Stderr returns a io.MultiWriter containing all
// plugins that implement Errer. If none are found,
// then os.Stderr is returned
func Stderr(plugs ...plugins.Plugin) io.Writer {
	var ins []io.Writer

	for _, p := range plugs {
		if x, ok := p.(Errer); ok {
			ins = append(ins, x.Stderr())
		}
	}

	if len(ins) == 0 {
		return os.Stderr
	}
	return io.MultiWriter(ins...)
}
