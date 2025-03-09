package fgcli

var globalOpts = &CommandOpts{}

type CommandOpts struct {
	version string
	help    string

	enableConfig bool
	cfgFile      string
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

func EnableConfig(enableConfig bool) Option {
	return func(o *CommandOpts) {
		o.enableConfig = enableConfig
	}
}

func SetCfgFile(cfgFile string) Option {
	return func(o *CommandOpts) {
		o.cfgFile = cfgFile
	}
}
