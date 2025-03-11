package fgcli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	commandUsage  = "[command] [flags]"
	flagArgsUsage = "[flags] [args]"
)

// Commander 定义 command 的接口
type Commander interface {
	// NameOrUsage 返回 command 或者用法
	// 如果返回值第二个参数为 true，则表示该 command 需要显示用法
	NameOrUsage() (string, bool)

	// Short 返回 command 的短描述 -- short
	Short() string

	// Long 返回 command 的描述 -- long
	Long() string

	// PreRun 定义 command 的初始化函数
	PreRun(ctx context.Context, args []string) error

	// Run 定义 command 的执行函数
	Run(ctx context.Context, args []string) error

	// Commands 返回 command 的子命令
	Commands() []Commander

	Flags() Flager
}

type rcommand struct {
	rcobra *cobra.Command
}

func (r *rcommand) applyOptions() error {
	if globalOpts.help != "" {
		r.rcobra.SetHelpTemplate(globalOpts.help)
	}

	if globalOpts.version != "" {
		r.rcobra.SetVersionTemplate(globalOpts.version)
	}

	if globalOpts.enableConfig || globalOpts.configFile != "" {
		bindConfigFlag(r, globalOpts.configFile)
		cobra.OnInitialize(setConfig())
		cobra.OnInitialize(readInConfig())
	}
	return nil
}

type commandBuilder struct {
	commander    Commander
	cobraCommand *cobra.Command

	commands []*commandBuilder
}

func (cb *commandBuilder) build() error {
	cb.cobraCommand = &cobra.Command{
		Use:   nameOrUsage(cb.commander),
		Short: cb.commander.Short(),
		Long:  cb.commander.Long(),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cb.commander.PreRun(cmd.Context(), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cb.commander.Run(cmd.Context(), args)
		},
		SilenceErrors:              true,
		SilenceUsage:               true,
		SuggestionsMinimumDistance: 2,
	}

	applyFlags(cb.cobraCommand, cb.commander.Flags())

	if globalOpts.enableConfig || globalOpts.configFile != "" {
		_ = bindViper(cb.cobraCommand)
	}

	for _, sub := range cb.commands {
		if err := sub.build(); err != nil {
			return err
		}
		cb.cobraCommand.AddCommand(sub.cobraCommand)
	}

	return nil
}

func bindConfigFlag(r *rcommand, file string) {
	r.rcobra.Flags().StringVarP(&file, "config", "c", file, "config file")
	_ = vc.v.BindPFlag("config", r.rcobra.Flags().Lookup("config"))
	// // r.rcobra.Flags().StringArrayVarP(&dirs, "configDir", "", dirs, "config dir")
	// // _ = viper.BindPFlag("configDir", r.rcobra.Flags().Lookup("configDir"))
}

func nameOrUsage(cmd Commander) string {
	use, ok := cmd.NameOrUsage()
	if ok {
		return use
	}
	var usage = commandUsage
	if cmd.Commands() == nil {
		usage = flagArgsUsage
	}
	return fmt.Sprintf("%s %s", use, usage)
}

// SimpleCommand 定义一个 command 实现了 Commander 接口
type SimpleCommand struct {
	// usage 定义 command 的用法
	// 如果 usage 不为空，则使用 usage 而不使用 CmdName
	Usage    string
	CmdName  string
	CmdShort string
	CmdLong  string
	Args     *cobra.PositionalArgs

	PreRunFunc func(ctx context.Context, args []string) error
	RunFunc    func(ctx context.Context, args []string) error
	// Withc func(cmd *cobra.Command, r *Command)
	// Initc func(cd *Commander)

	Commanders []Commander
	Flager     Flager
}

func (c *SimpleCommand) NameOrUsage() (string, bool) {
	if c.Usage != "" {
		return c.Usage, true
	}
	return c.CmdName, false
}

func (c *SimpleCommand) Short() string {
	return c.CmdShort
}

func (c *SimpleCommand) Long() string {
	return c.CmdLong
}

func (c *SimpleCommand) PreRun(ctx context.Context, args []string) error {
	if c.PreRunFunc != nil {
		return c.PreRunFunc(ctx, args)
	}
	return nil
}

func (c *SimpleCommand) Run(ctx context.Context, args []string) error {
	if globalOpts.enableConfig || globalOpts.configFile != "" {
		printConfig()
		if err := vc.v.Unmarshal(c.Flager); err != nil {
			return err
		}
	}
	if c.RunFunc != nil {
		return c.RunFunc(ctx, args)
	}
	return nil
}

func (c *SimpleCommand) Commands() []Commander {
	return c.Commanders
}

func (c *SimpleCommand) Flags() Flager {
	return c.Flager
}
