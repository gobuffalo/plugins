package plugins

// Plugin is the most basic interface a plugin can implement.
type Plugin interface {
	PluginName() string
}

// Scoper can be implemented to return a slice of plugins that
// are important to the type defining it.
type Scoper interface {
	ScopedPlugins() []Plugin
}

type Feeder func() []Plugin

// Needer can be implemented to receive a Feeder function
// that can be used to gain access to other plugins in the system.
type Needer interface {
	WithPlugins(Feeder)
}
