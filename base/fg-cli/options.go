package fgcli

type CommandOpts struct {
	usage string
	name  string
	short string
	long  string

	version string
	help    string

	cfgFile string
	cfgDir  []string
}

type Option func(*CommandOpts)

func WithName(name string) Option {
	return func(o *CommandOpts) {
		o.name = name
	}
}

func WithUsage(usage string) Option {
	return func(o *CommandOpts) {
		o.usage = usage
	}
}

func WithShort(short string) Option {
	return func(o *CommandOpts) {
		o.short = short
	}
}

func WithLong(long string) Option {
	return func(o *CommandOpts) {
		o.long = long
	}
}

func WithVersion(version string) Option {
	return func(o *CommandOpts) {
		o.version = version
	}
}

func WithHelp(help string) Option {
	return func(o *CommandOpts) {
		o.help = help
	}
}

func WithCfgFile(cfgFile string) Option {
	return func(o *CommandOpts) {
		o.cfgFile = cfgFile
	}
}

func WithCfgDir(cfgDir ...string) Option {
	return func(o *CommandOpts) {
		o.cfgDir = cfgDir
	}
}
