package main

import (
	cli "github.com/chhz0/gokit/pkg/cli"
	"github.com/spf13/pflag"
)

type RootOption struct {
	AppName string `mapstructure:"app_name" yaml:"app_name"`
	Version string `mapstructure:"version" yaml:"version"`
}

// LocalFlags implements cli.Flager.
func (r *RootOption) LocalFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	pfs.StringVarP(&r.Version, "version", "v", "", "app version")
	pfs.StringVarP(&r.AppName, "app_name", "a", "", "app name")

	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

// PersistentFlags implements cli.Flager.
func (r *RootOption) PersistentFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

var _ cli.Flager = (*RootOption)(nil)

type PrintOption struct {
	Print string `mapstructure:"print"`
	From  string `mapstructure:"from"`
}

func (p *PrintOption) LocalFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	pfs.StringVarP(&p.Print, "print", "p", "print", "print")
	pfs.StringVarP(&p.From, "from", "f", "from", "from")

	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (p *PrintOption) PersistentFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

type EchoOption struct {
	Echo string      `mapstructure:"echo"`
	Time TimesOption `mapstructure:"time"`
}

func (e *EchoOption) LocalFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	pfs.StringVarP(&e.Echo, "echo", "e", "echo", "echo")
	pfs.AddFlagSet(e.Time.LocalFlags(pfs).PFlags)

	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (e *EchoOption) PersistentFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

type TimesOption struct {
	Time int `mapstructure:"times"`
}

func (t *TimesOption) LocalFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	pfs.IntVarP(&t.Time, "times", "t", 5, "times")

	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}

func (t *TimesOption) PersistentFlags(pfs *pflag.FlagSet) *cli.FlagSet {
	return &cli.FlagSet{
		PFlags:   pfs,
		Required: []string{},
	}
}
