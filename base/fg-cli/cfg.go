package fgcli

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	defaultDir        = ".fgcli"
	defaultConfigFile = "config.yaml"
)

func initConfigFunc(file, dir string) {
	if file != "" {
		viper.AddConfigPath(dir)
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(defaultDir)

		// Find home directory.
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
		_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", file, err)
		os.Exit(1)
	}
}
