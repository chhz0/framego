package main

import (
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
	pfs.StringVarP(&r.Version, "version", "v", "v0.0.1", "app version")
	pfs.StringVarP(&r.AppName, "app", "a", "fg-cli", "app name")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

// PersistentFlags implements fgcli.Flager.
func (r *RootOption) PersistentFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("root", pflag.ExitOnError)
	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

var _ fgcli.Flager = (*RootOption)(nil)

type PrintOption struct {
	Print string `mapstructure:"print"`
	From  string `mapstructure:"from"`
}

func (p *PrintOption) LocalFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("print", pflag.ExitOnError)
	pfs.StringVarP(&p.Print, "print", "p", "print", "print")
	pfs.StringVarP(&p.From, "from", "f", "from", "from")

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
	Time TimesOption `mapstructure:"time"`
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
	Time int `mapstructure:"times"`
}

func (t *TimesOption) LocalFlags() *fgcli.FlagSet {
	pfs := pflag.NewFlagSet("times", pflag.ExitOnError)
	pfs.IntVarP(&t.Time, "times", "t", 5, "times")

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
