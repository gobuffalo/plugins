package plugprint

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/gobuffalo/plugins"
)

// Print will try and print a helpful Usage printing
// of the plugin and any plugins that are provided.
// 	$ buffalo
// 	---------
//
// 	Usage of buffalo:
// 	-h, --help   print this help
//
// 	Available Commands:
// 		buffalo fix      Attempt to fix a Buffalo application's API to match version in go.mod
// 		buffalo info     Print diagnostic information (useful for debugging)
// 		buffalo version  Print the version information
func Print(w io.Writer, main plugins.Plugin) error {
	if d, ok := main.(Describer); ok {
		fmt.Fprintf(w, "%s\n\n", d.Description())
	}

	header := strings.TrimSpace(fmt.Sprintf("%s", stringer(main)))
	header = fmt.Sprintf("$ %s", header)
	fmt.Fprintln(w, header)
	for i := 0; i < len(header); i++ {
		fmt.Fprint(w, "-")
	}
	fmt.Fprintln(w)

	if a, ok := main.(Aliaser); ok {
		aliases := a.CmdAliases()
		if len(aliases) != 0 {
			const al = "\nAliases:\n"
			fmt.Fprint(w, al)
			fmt.Fprintln(w, strings.Join(aliases, ", "))
		}
	}

	if u, ok := main.(UsagePrinter); ok {
		fmt.Fprintln(w)
		if err := u.PrintUsage(w); err != nil {
			return err
		}
	}

	if u, ok := main.(FlagPrinter); ok {
		const fp = "\nFlags:\n"
		fmt.Fprint(w, fp)
		if err := u.PrintFlags(w); err != nil {
			return err
		}
	}

	if err := printCommands(w, main); err != nil {
		return err
	}

	if err := printPlugins(w, main); err != nil {
		return err
	}

	return nil
}

func printPlugins(w io.Writer, main plugins.Plugin) error {
	mm := map[string]plugins.Plugin{}
	usingPlugins(main, mm)
	if len(mm) == 0 {
		return nil
	}

	plugs := make([]plugins.Plugin, 0, len(mm))
	for _, p := range mm {
		plugs = append(plugs, p)
	}
	sort.Slice(plugs, func(i, j int) bool {
		return plugs[i].PluginName() < plugs[j].PluginName()
	})

	fmt.Fprintln(w, "\nUsing Plugins:")
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "\t%s\t%s\t%s\n", "Name", "Description", "Type")
	fmt.Fprintf(tw, "\t%s\t%s\t%s\n", "----", "-----------", "----")
	for _, p := range plugs {
		if _, ok := p.(Hider); ok {
			continue
		}
		fmt.Fprintf(tw, "\t%s\t%s\t%s\n", p.PluginName(), desc(p), typeName(p))
	}

	tw.Flush()
	return nil
}

func typeName(p plugins.Plugin) string {
	rv := reflect.Indirect(reflect.ValueOf(p))
	rt := reflect.TypeOf(rv.Interface())

	bb := &bytes.Buffer{}

	t := fmt.Sprintf("%T", p)
	if strings.HasPrefix(t, "*") {
		fmt.Fprint(bb, "*")
		t = strings.TrimPrefix(t, "*")
	}
	fmt.Fprintf(bb, "%s/", path.Dir(rt.PkgPath()))
	fmt.Fprint(bb, t)

	return bb.String()
}

func usingPlugins(plug plugins.Plugin, mm map[string]plugins.Plugin) {
	if _, ok := mm[plug.PluginName()]; ok {
		return
	}
	mm[plug.PluginName()] = plug

	wp, ok := plug.(plugins.Scoper)
	if !ok {
		return
	}

	for _, p := range wp.ScopedPlugins() {
		mm[p.PluginName()] = p
	}

	return
}

func printCommands(w io.Writer, main plugins.Plugin) error {
	sc, ok := main.(SubCommander)
	if !ok {
		return nil
	}

	plugs := sc.SubCommands()
	if len(plugs) == 0 {
		return nil
	}

	sort.Slice(plugs, func(i, j int) bool {
		return plugs[i].PluginName() < plugs[j].PluginName()
	})

	const ac = "\nAvailable Commands:\n"
	fmt.Fprint(w, ac)

	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "\t%s\t%s\n", "Command", "Description")
	fmt.Fprintf(tw, "\t%s\t%s\n", "-------", "-----------")
	for _, c := range plugs {
		line := fmt.Sprintf("\t%s\t%s", cmdName(c), desc(c))
		fmt.Fprintln(tw, line)
	}
	tw.Flush()
	return nil
}

func stringer(p plugins.Plugin) string {
	if s, ok := p.(fmt.Stringer); ok {
		return s.String()
	}
	return cmdName(p)
}

func desc(p plugins.Plugin) string {
	if d, ok := p.(Describer); ok {
		return d.Description()
	}
	return ""
}

func cmdName(p plugins.Plugin) string {
	if d, ok := p.(CommandNamer); ok {
		return d.CmdName()
	}
	return path.Base(p.PluginName())
}
