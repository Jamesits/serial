package main

import (
	"github.com/Jamesits/serial/internal/cmd/list"
	"github.com/Jamesits/serial/internal/cmd/open"
	"github.com/Jamesits/serial/internal/cmd/root"
	"github.com/Jamesits/serial/internal/cmd/version"
	"github.com/Jamesits/serial/pkg/logging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func main() {
	defer quitHook()
	logging.InitGlobalLogger(quitHook)

	log.Traceln("init cobra")
	cobra.OnInitialize(initCobraConfig)
	root.CommandDefinition.SilenceErrors = true

	log.Traceln("init subcommands")
	root.CommandDefinition.AddCommand(version.CommandDefinition)
	root.CommandDefinition.AddCommand(list.CommandDefinition)
	root.CommandDefinition.AddCommand(open.CommandDefinition)

	log.Traceln("cobra entry")
	if err := root.CommandDefinition.Execute(); err != nil {
		//log.WithError(err).Errorln("error")
		os.Exit(1)
	}
}

// quitHook runs when the whole program is quitting
func quitHook() {
	log.Traceln("quitHook() start")
}

func initCobraConfig() {
	if root.UserConfigFilePath != "" {
		// Use config file from the flag.
		log.Traceln("read user-supplied config file")
		viper.SetConfigFile(root.UserConfigFilePath)
	} else {
		log.Traceln("read default config file")

		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".serial")
	}

	log.Traceln("init viper")
	viper.AutomaticEnv()

	log.Traceln("viper read config")
	if err := viper.ReadInConfig(); err == nil {
		log.Traceln("using config file:", viper.ConfigFileUsed())
	} else {
		log.Traceln("config file parsing failed")
	}
}
