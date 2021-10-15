package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var CommandDefinition = &cobra.Command{
	Use:   os.Args[0],
	Short: "Serial port connector",
	Long: `A simple program to interact with serial ports that just works.
Go to https://github.com/Jamesits/serial for detailed documentation and source code.
`,
	PersistentPreRun: GlobalSetup,
	RunE:             main,
}

var UserConfigFilePath string
var LogLevel uint8

func init() {
	CommandDefinition.PersistentFlags().StringVar(&UserConfigFilePath, "config", "", "config file (default is $HOME/.cobra.yaml)")
	CommandDefinition.PersistentFlags().Uint8VarP(&LogLevel, "loglevel", "v", uint8(log.InfoLevel), "Log verbose level (0-6)")
}

func GlobalSetup(cmd *cobra.Command, args []string) {
	log.SetLevel(log.Level(LogLevel))
}

func main(cmd *cobra.Command, args []string) error {
	log.Errorf("please supply appropriate commands; run \"%s --help\" to get a list of commands available", os.Args[0])

	return nil
}
