package gokit

import "github.com/chhz0/gokit/pkg/cli"

func NewCli(cmd cli.Commander, opts ...cli.Option) (cli.CliExector, error) {
	return cli.New(cmd, opts...)
}

func Options() {
	
}