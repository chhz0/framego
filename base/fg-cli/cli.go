package fgcli

import (
	"context"
)

type CliExector interface {
	Execute(ctx context.Context) error
}

type Cli struct {
	*rcommand
}

func (c *Cli) Execute(ctx context.Context) error {
	return c.rcobra.ExecuteContext(ctx)
}

func NewCli(cmd Commander, opts ...Option) (CliExector, error) {
	for _, opt := range opts {
		opt(globalOpts)
	}

	r, err := compile(cmd)
	if err != nil {
		return nil, err
	}
	_ = r.applyOptions()
	return &Cli{r}, nil
}

func compile(rcmd Commander) (*rcommand, error) {
	rbuilder := &commandBuilder{
		commander: rcmd,
	}

	var addCmdBuilder func(cb *commandBuilder, cmder Commander)
	addCmdBuilder = func(cb *commandBuilder, cmder Commander) {
		cb2 := &commandBuilder{
			commander: cmder,
		}
		cb.commands = append(cb.commands, cb2)
		for _, c := range cmder.Commands() {
			addCmdBuilder(cb2, c)
		}
	}

	for _, cmder := range rcmd.Commands() {
		addCmdBuilder(rbuilder, cmder)
	}

	if err := rbuilder.build(); err != nil {
		return nil, err
	}

	return &rcommand{rcobra: rbuilder.cobraCommand}, nil
}
