package fgcli

import "github.com/spf13/viper"

var globalOpts = &CommandOpts{}

type CommandOpts struct {
	version string
	help    string

	enableConfig  bool
	configFile    string
	configHandler func()
}

type Option func(*CommandOpts)

func SetVersion(version string) Option {
	return func(o *CommandOpts) {
		o.version = version
	}
}

func SetHelp(help string) Option {
	return func(o *CommandOpts) {
		o.help = help
	}
}

func EnableConfig(v *viper.Viper) Option {
	if v == nil {
		v = viper.New()
	}
	vc = &vconfig{v}

	return func(o *CommandOpts) {
		o.enableConfig = true
	}
}

func SetCfgFile(cfgFile string) Option {
	return func(o *CommandOpts) {
		o.configFile = cfgFile
	}
}

func SetConfigHandler(fn func()) Option {
	return func(o *CommandOpts) {
		o.configHandler = fn
	}
}
