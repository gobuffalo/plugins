package plugio

import (
	"io"
	"os"

	"github.com/gobuffalo/plugins"
)

var _ Inner = inner{}

type inner struct {
	io.Reader
}

func (i inner) PluginName() string {
	return "inner"
}

func (i inner) Stdin() io.Reader {
	return i.Reader
}

// NewInner converts an io.Reader to an Inner plugin.
func NewInner(r io.Reader) Inner {
	return inner{Reader: r}
}

// Stdin returns a io.MultiReader containing all
// plugins that implement Inner. If none are found,
// then os.Stdin is returned
func Stdin(plugs ...plugins.Plugin) io.Reader {
	var ins []io.Reader

	for _, p := range plugs {
		if x, ok := p.(Inner); ok {
			ins = append(ins, x.Stdin())
		}
	}

	if len(ins) == 0 {
		return os.Stdin
	}
	return io.MultiReader(ins...)
}
