package main

import (
	"context"
	"fmt"
	"strings"

	fgcli "github.com/chhz0/framego/base/fg-cli"
	fgconfig "github.com/chhz0/framego/base/fg-config"
)

func newPrintCmd() *fgcli.SimpleCommand {
	opt := &PrintOption{}
	return &fgcli.SimpleCommand{
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
func newEchoCmd() *fgcli.SimpleCommand {
	opt := &EchoOption{}
	return &fgcli.SimpleCommand{
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
		Commanders: []fgcli.Commander{
			newTimeCmd(),
		},
	}
}

func newTimeCmd() *fgcli.SimpleCommand {
	opt := &TimesOption{}
	return &fgcli.SimpleCommand{
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

func testfgConfig() *fgconfig.VConfig {
	vc := fgconfig.NewWith(
		fgconfig.WithConfig(&fgconfig.LocalConfig{
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
	cli, err := fgcli.NewCli(&fgcli.SimpleCommand{
		CmdName:  "fgcli",
		CmdShort: "create cli application",
		CmdLong:  "fg-cli is quickly create command line application's framework",
		PreRunFunc: func(ctx context.Context, args []string) error {
			fmt.Println("fg-cli pre run ...")
			return nil
		},
		RunFunc: func(ctx context.Context, args []string) error {
			fmt.Println("fg-cli run ...")
			fmt.Printf("fg-cli options: %v\n", opt)
			return nil
		},
		Flager: opt,
		Commanders: []fgcli.Commander{
			newPrintCmd(),
			newEchoCmd(),
		},
	},
		// --- use fgcli config
		// fgcli.EnableConfig(nil),
		// fgcli.SetCfgFile("./config/confg.yaml"),
		// --- test fgconfig
		fgcli.EnableConfig(v.V()),
		fgcli.SetConfigHandler(v.Load),
	)
	if err != nil {
		panic(err)
	}

	if err := cli.Execute(context.Background()); err != nil {
		panic(err)
	}
}
