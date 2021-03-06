/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/conreality/conreality.go/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const masterURL = "localhost:50051"
const playerNick = "Admin"
const timeout = 1 * time.Second

var configFile string
var debug bool
var verbose bool

// RootCmd describes the `concli` command
var RootCmd = &cobra.Command{
	Use:     "concli",
	Short:   "Conreality Command-Line Interface (CLI)",
	Version: sdk.Version,
}

// Execute implements the `concli` command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "C", "", "Set config file (default: $HOME/.conreality/config.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debugging")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	RootCmd.SetVersionTemplate(`Conreality CLI {{printf "%s" .Version}}
`)
}

// initConfig reads the configuration file and environment variables.
func initConfig() {
	if configFile != "" {
		// Use the specified config file:
		viper.SetConfigFile(configFile)
	} else {
		// Search for config file in the current directory and under the home directory:
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.conreality")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in:
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
