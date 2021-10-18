package main

import (
	"github.com/Jamesits/serial/internal/cmd/cu"
	"github.com/Jamesits/serial/pkg/logging"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	defer quitHook()
	logging.InitGlobalLogger(quitHook)
	cobra.OnInitialize()
	cu.CommandDefinition.SilenceErrors = true

	if err := cu.CommandDefinition.Execute(); err != nil {
		//log.WithError(err).Errorln("error")
		os.Exit(1)
	}
}

// quitHook runs when the whole program is quitting
func quitHook() {

}
