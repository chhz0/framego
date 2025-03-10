package fgcli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Flager interface {
	PersistentFlags() *FlagSet
	LocalFlags() *FlagSet
}

type FlagSet struct {
	PFlags   *pflag.FlagSet
	Required []string
}

func applyFlags(cmd *cobra.Command, flager Flager) {
	if cmd == nil || flager == nil {
		return
	}

	applyLocal(cmd, flager.LocalFlags())
	applyPeristent(cmd, flager.PersistentFlags())
}

func applyLocal(cmd *cobra.Command, fs *FlagSet) {
	cmd.Flags().AddFlagSet(fs.PFlags)
	for _, rflag := range fs.Required {
		_ = cmd.MarkFlagRequired(rflag)
	}
}

func applyPeristent(cmd *cobra.Command, fs *FlagSet) {
	cmd.PersistentFlags().AddFlagSet(fs.PFlags)
	for _, rflag := range fs.Required {
		_ = cmd.MarkPersistentFlagRequired(rflag)
	}
}
