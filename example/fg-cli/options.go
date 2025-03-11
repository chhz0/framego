package main

import (
	fgcli "github.com/chhz0/framego/base/fg-cli"
	"github.com/spf13/pflag"
)

type RootOption struct {
	AppName string `mapstructure:"app_name" yaml:"app_name"`
	Version string `mapstructure:"version" yaml:"version"`
}

// LocalFlags implements fgcli.Flager.
func (r *RootOption) LocalFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	pfs.StringVarP(&r.Version, "version", "v", "", "app version")
	pfs.StringVarP(&r.AppName, "app_name", "a", "", "app name")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

// PersistentFlags implements fgcli.Flager.
func (r *RootOption) PersistentFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
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

func (p *PrintOption) LocalFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	pfs.StringVarP(&p.Print, "print", "p", "print", "print")
	pfs.StringVarP(&p.From, "from", "f", "from", "from")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (p *PrintOption) PersistentFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

type EchoOption struct {
	Echo string      `mapstructure:"echo"`
	Time TimesOption `mapstructure:"time"`
}

func (e *EchoOption) LocalFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	pfs.StringVarP(&e.Echo, "echo", "e", "echo", "echo")
	pfs.AddFlagSet(e.Time.LocalFlags(pfs).PFlags)

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (e *EchoOption) PersistentFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

type TimesOption struct {
	Time int `mapstructure:"times"`
}

func (t *TimesOption) LocalFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	pfs.IntVarP(&t.Time, "times", "t", 5, "times")

	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (t *TimesOption) PersistentFlags(pfs *pflag.FlagSet) *fgcli.FlagSet {
	return &fgcli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}
