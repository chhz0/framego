package main

import (
	"time"

	fgcli "github.com/chhz0/framego/base/fg-cli"
	"github.com/spf13/pflag"
)

type RootOption struct {
	AppName string `mapstructure:"app_name"`
	Version string `mapstructure:"version"`
}

// LocalFlags implements fgcli.Flager.
func (r *RootOption) LocalFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("root", pflag.ExitOnError)
	pfs.StringVarP(&r.Version, "version", "v", "v1.0.0", "app version")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

// PersistentFlags implements fgcli.Flager.
func (r *RootOption) PersistentFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("root", pflag.ExitOnError)
	pfs.StringVarP(&r.AppName, "app-name", "a", "fg-cli", "app name")
	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{"app-name"},
	}
}

var _ fgcli.Flager = (*RootOption)(nil)

type PrintOption struct {
	print string `mapstructure:"print"`
	from  string `mapstructure:"from"`
}

func (p *PrintOption) LocalFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("print", pflag.ExitOnError)
	pfs.StringVarP(&p.print, "print", "p", "print", "print")
	pfs.StringVarP(&p.from, "from", "f", "from", "from")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (p *PrintOption) PersistentFlags() *fgcli.FlagSet {
	return &fgcli.FlagSet{
		PFlags:   nil,
		Required: []string{},
	}
}

type EchoOption struct {
	Echo string      `mapstructure:"echo"`
	Time TimesOption `mapstructure:"timeOption"`
}

func (e *EchoOption) LocalFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("echo", pflag.ExitOnError)
	pfs.StringVarP(&e.Echo, "echo", "e", "echo", "echo")

	pfs.AddFlagSet(e.Time.LocalFlags().PFlags)

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (e *EchoOption) PersistentFlags() *fgcli.FlagSet {
	return &fgcli.FlagSet{
		PFlags:   nil,
		Required: []string{},
	}
}

type TimesOption struct {
	Time time.Duration `mapstructure:"times"`
}

func (t *TimesOption) LocalFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("times", pflag.ExitOnError)
	pfs.DurationVarP(&t.Time, "times", "t", time.Second, "times")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (t *TimesOption) PersistentFlags() *fgcli.FlagSet {
	return &fgcli.FlagSet{
		PFlags:   nil,
		Required: []string{},
	}
}
