package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/chhz0/gokit"
	"github.com/chhz0/gokit/pkg/cli"
	"github.com/chhz0/gokit/pkg/config"
)

func newPrintCmd() *cli.SimpleCommand {
	opt := &PrintOption{}
	return &cli.SimpleCommand{
		CmdName:  "print",
		CmdShort: "Print anything to the screen",
		CmdLong: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		RunFunc: func(ctx context.Context, args []string) error {
			fmt.Println("Print: " + opt.Print)
			fmt.Println("From: " + opt.From)
			return nil
		},
		Flager: opt,
	}
}
func newEchoCmd() *cli.SimpleCommand {
	opt := &EchoOption{}
	return &cli.SimpleCommand{
		CmdName:  "echo",
		CmdShort: "Echo anything to the screen",
		CmdLong: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		RunFunc: func(ctx context.Context, args []string) error {
			fmt.Println("Echo: " + strings.Join(args, " "))
			fmt.Printf("echo options: %v\n", opt)
			return nil
		},
		Flager: opt,
		Commanders: []cli.Commander{
			newTimeCmd(),
		},
	}
}

func newTimeCmd() *cli.SimpleCommand {
	opt := &TimesOption{}
	return &cli.SimpleCommand{
		Usage:    "times [# times] [string to echo]",
		CmdShort: "Echo anything to the screen more times",
		CmdLong: `echo things multiple times back to the user by providing
a count and a string.`,
		RunFunc: func(ctx context.Context, args []string) error {
			for i := 0; i < opt.Time; i++ {
				fmt.Println("Echo times: " + strings.Join(args, " "))
			}
			return nil
		},
		Flager: opt,
	}
}

func testfgConfig() *config.VConfig {
	vc := config.NewWith(
		config.WithConfig(&config.LocalConfig{
			ConfigName:  "test",
			ConfigType:  "yaml",
			ConfigPaths: []string{"./config"},
		}),
	)
	return vc
}

func main() {
	v := testfgConfig()
	opt := &RootOption{}
	cli, err := gokit.NewCli(
		&cli.SimpleCommand{
			CmdName:  "gkcli",
			CmdShort: "create cli application",
			CmdLong:  "gk-cli is quickly create command line application's framework",
			PreRunFunc: func(ctx context.Context, args []string) error {
				fmt.Println("gk-cli pre run ...")
				return nil
			},
			RunFunc: func(ctx context.Context, args []string) error {
				fmt.Println("gk-cli run ...")
				fmt.Printf("gk-cli options: %v\n", opt)
				return nil
			},
			Flager: opt,
			Commanders: []cli.Commander{
				newPrintCmd(),
				newEchoCmd(),
			},
		},
		// --- use cli config
		// cli.EnableConfig(nil),
		// cli.SetCfgFile("./config/confg.yaml"),
		// --- test fgconfig
		cli.EnableConfig(v.V()),
		cli.SetConfigHandler(v.Load),
	)
	if err != nil {
		panic(err)
	}

	if err := cli.Execute(context.Background()); err != nil {
		panic(err)
	}
}
