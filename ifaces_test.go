package plugins

// var _ Plugin = aliaser{}
var _ Plugin = back("")

// var _ Plugin = namer("")

// var _ cmdAliaser = aliaser{}
// var _ cmdNamer = namer("")

type back string

func (b back) PluginName() string {
	return string(b)
}

// type aliaser struct{}
//
// func (a aliaser) PluginName() string {
// 	return "aliaser"
// }
//
// func (a aliaser) CmdAliases() []string {
// 	return []string{"b", "be"}
// }
//
// type namer string
//
// func (a namer) PluginName() string {
// 	return "namer"
// }
//
// func (a namer) CmdName() string {
// 	return string(a)
// }
