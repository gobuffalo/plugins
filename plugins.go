package plugins

// Plugins is a slice of type `Plugin` that provides
// additional useful functionality.
type Plugins []Plugin

// ScopedPlugins implements Scoper, return itself.
func (plugs Plugins) ScopedPlugins() []Plugin {
	return plugs
}
