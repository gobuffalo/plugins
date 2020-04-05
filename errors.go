package plugins

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"strings"
)

func Wrap(p Plugin, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s (%s): %w", typeName(p), p.PluginName(), err)
}

func typeName(p Plugin) string {
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
