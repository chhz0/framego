package fgcli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultDir        = ".fgcli"
	defaultConfigFile = "config.yaml"
)

var vc *vconfig

type vconfig struct {
	v *viper.Viper
}

func bindViper(cmd *cobra.Command) error {
	if err := vc.v.BindPFlags(cmd.PersistentFlags()); err != nil {
		return err
	}
	if err := vc.v.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	return nil
}

func setConfig() func() {
	return func() {
		flagConfig := vc.v.GetString("config")
		if flagConfig != "" {
			vc.v.SetConfigFile(flagConfig)
			path, file := filepath.Split(flagConfig)
			vc.v.SetConfigName(file)
			vc.v.AddConfigPath(path)
		}

		if globalOpts.configHandler != nil {
			cobra.OnInitialize(globalOpts.configHandler)
			return
		}

		if globalOpts.configFile != "" {
			vc.v.SetConfigFile(globalOpts.configFile)
			return
		}

		home, err := homedir.Dir()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read home dir(%s): %v\n", home, err)
			os.Exit(1)
		}

		configPath := filepath.Join(home, defaultConfigFile)
		vc.v.AddConfigPath(configPath)
		vc.v.SetConfigName(defaultConfigFile)
	}
}

func readInConfig() func() {
	return func() {
		if err := vc.v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				e := fmt.Errorf("Error: config file(%s) not found: %v\n", vc.v.ConfigFileUsed(), err)
				panic(e)
			}
			e := fmt.Errorf("Error: failed to read configuration file(%s): %v\n", vc.v.ConfigFileUsed(), err)
			panic(e)
		}
	}
}

func printConfig() {
	keys := vc.v.AllKeys()
	if len(keys) == 0 {
		return
	}
	fmt.Printf("\n%v Configuration items:\n", color.GreenString("✔"))
	table := uitable.New()
	table.AddRow("KEY", "VALUE")
	for _, k := range keys {
		table.AddRow(k, vc.v.Get(k))
	}
	fmt.Printf("%v\n\n", table)
}
