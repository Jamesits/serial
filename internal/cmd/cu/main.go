package cu

import (
	"github.com/spf13/cobra"
	"os"
)

var CommandDefinition = &cobra.Command{
	Use:              os.Args[0],
	Short:            "Call up another system",
	Long:             "Call up another system.",
	PersistentPreRun: GlobalSetup,
	RunE:             main,
}

var ParityFlagEven bool
var ParityFlagOdd bool
var ParityString string
var HalfDuplex bool
var NoStop bool
var NoHardwareFloeControl bool
var EscapeChar string

//var DstSystem string
//var DstPhone string
//var DstPort string
//var DstPortAlt string
var DstLine string
var Speed uint64

var PrintVersion bool

func init() {
	CommandDefinition.PersistentFlags().BoolVarP(&ParityFlagEven, "", "e", false, "Use even parity.")
	CommandDefinition.PersistentFlags().BoolVarP(&ParityFlagOdd, "", "o", false, "Use odd parity.")
	CommandDefinition.PersistentFlags().StringVar(&ParityString, "parity", "none", "Use the specified parity.")
	CommandDefinition.PersistentFlags().BoolVarP(&HalfDuplex, "halfduplex", "h", false, "Echo characters locally (half-duplex mode).")
	CommandDefinition.PersistentFlags().BoolVar(&NoStop, "nostop", false, "Turn off XON/XOFF handling (it is on by default).")
	CommandDefinition.PersistentFlags().BoolVarP(&NoHardwareFloeControl, "nortscts", "f", false, "Do not use hardware flow control.")
	CommandDefinition.PersistentFlags().StringVarP(&EscapeChar, "escape", "E", "~", "Set the escape character. To eliminate the escape character, use -E ''.")
	//CommandDefinition.PersistentFlags().StringVarP(&DstSystem, "system", "z", "", "The system to call.")
	//CommandDefinition.PersistentFlags().StringVarP(&DstPhone, "phone", "c", "", "The phone number to call.")
	//CommandDefinition.PersistentFlags().StringVarP(&DstPort, "port", "p", "", "Name the port to use.")
	//CommandDefinition.PersistentFlags().StringVarP(&DstPortAlt, "", "a", "", "Equivalent to --port port.")
	CommandDefinition.PersistentFlags().StringVarP(&DstLine, "line", "l", "", "Name the line to use by giving a device name.  This may be used to dial out on ports that are not listed in the UUCP configuration files.  Write access to the device is required.")
	CommandDefinition.PersistentFlags().Uint64VarP(&Speed, "speed", "s", 9600, "The speed (baud rate) to use.")

	CommandDefinition.PersistentFlags().BoolVarP(&PrintVersion, "version", "v", false, "Report version information and exit.")
}

func GlobalSetup(cmd *cobra.Command, args []string) {

}

func main(cmd *cobra.Command, args []string) error {

	return nil
}
