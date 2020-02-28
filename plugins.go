package plugins

import "sort"

// Plugins is a slice of type `Plugin` that provides
// additional useful functionality.
type Plugins []Plugin

// Len is the number of elements in the collection.
func (p Plugins) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p Plugins) Less(i int, j int) bool {
	return p[i].PluginName() < p[j].PluginName()
}

// Swap swaps the elements with indexes i and j.
func (p Plugins) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// ScopedPlugins implements Scoper, return itself.
func (plugs Plugins) ScopedPlugins() []Plugin {
	return plugs
}

// Sort will sort a slice of plugins, by wrapping them
// in Plugins and using it's sort.Interface implementation.
func Sort(plugs []Plugin) {
	sort.Sort(Plugins(plugs))
}
