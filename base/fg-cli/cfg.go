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

func bindConfigFlag(r *rcommand, file string) {
	r.rcobra.Flags().StringVarP(&file, "config", "c", file, "config file")
	_ = viper.BindPFlag("config", r.rcobra.Flags().Lookup("config"))
	// // r.rcobra.Flags().StringArrayVarP(&dirs, "configDir", "", dirs, "config dir")
	// // _ = viper.BindPFlag("configDir", r.rcobra.Flags().Lookup("configDir"))
}

func bindViper(cmd *cobra.Command) error {
	if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
		return err
	}
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	return nil
}

func useConfig(file string) func() {
	return func() {
		if file != "" {
			viper.SetConfigFile(file)
			flagConfig := viper.GetString("config")
			if flagConfig != file {
				viper.SetConfigFile(flagConfig)
			}
		} else {
			home, err := homedir.Dir()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read home dir(%s): %v\n", home, err)
				os.Exit(1)
			}

			configPath := filepath.Join(home, defaultConfigFile)
			viper.AddConfigPath(configPath)
			viper.SetConfigName(defaultConfigFile)
		}

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				e := fmt.Errorf("Error: config file(%s) not found: %v\n", viper.ConfigFileUsed(), err)
				panic(e)
			}
			e := fmt.Errorf("Error: failed to read configuration file(%s): %v\n", viper.ConfigFileUsed(), err)
			panic(e)
		}

	}
}

func printConfig() {
	keys := viper.AllKeys()
	if len(keys) == 0 {
		return
	}
	fmt.Printf("\n%v Configuration items:\n", color.GreenString("✔"))
	table := uitable.New()
	table.AddRow("KEY", "VALUE")
	for _, k := range keys {
		table.AddRow(k, viper.Get(k))
	}
	fmt.Printf("%v\n\n", table)
}
