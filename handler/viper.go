package handler

import (
	"github.com/materials-commons/config"
)

const (
	// Use the default handler
	Default = "Default"

	// Use the Override handler
	Override = "Override"

	// Use the environment handler
	Environment = "Environment"

	// Use the ini handler
	Ini = "Ini"
)

// Viper implements github.com/spf13/viper
func Viper(loader config.Loader) config.Handler {
	return LowercaseKey(
		Prioritized(
			NameHandler(Default, Map()),
			NameHandler(Override, Map()),
			NameHandler(Ini, Loader(loader))))
}

// ViperCaseSensitive implements github.com/spf13/viper except that keys
// are case sensitive.
func ViperCaseSensitive(loader config.Loader) config.Handler {
	return Prioritized(
		NameHandler(Default, Map()),
		NameHandler(Override, Map()),
		NameHandler(Ini, Loader(loader)))
}

// ViperEx implements github.com/spf13/viper with the addition of environment
// variables checked before checking for values in the ini file(s).
func ViperEx(loader config.Loader) config.Handler {
	return LowercaseKey(
		Prioritized(
			NameHandler(Default, Map()),
			NameHandler(Override, Map()),
			NameHandler(Environment, Env()),
			NameHandler(Ini, Loader(loader))))
}

// ViperExCaseSensitive implements ViperEx except that keys are case sensitive.
func ViperExCaseSensitive(loader config.Loader) config.Handler {
	return Prioritized(
		NameHandler(Default, Map()),
		NameHandler(Override, Map()),
		NameHandler(Environment, Env()),
		NameHandler(Ini, Loader(loader)))
}
