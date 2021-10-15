package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configVersionFull bool

var CommandDefinition = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		if configVersionFull {
			fmt.Println(getVersionFullString())
		} else {
			fmt.Println(getVersionNumberString())
		}
	},
}

func init() {
	CommandDefinition.PersistentFlags().BoolVar(&configVersionFull, "full", false, "Print a full version string in user-agent format")
}
