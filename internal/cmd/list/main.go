package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jszwec/csvutil"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.bug.st/serial"
	"strings"
)

var CommandDefinition = &cobra.Command{
	Use:   "list",
	Short: "List all serial ports",
	Long:  `List all serial ports available on the system, in various forms.`,
	RunE:  main,
}

var configPrint0 bool // print with "\0" as the separator; useful when running with Linux philosophy
var configPrintFormat string
var configPrintSerializeWrapInObject bool

func init() {
	CommandDefinition.PersistentFlags().BoolVar(&configPrint0, "print0", false, "Use '\\0' as the separator")
	CommandDefinition.PersistentFlags().StringVar(&configPrintFormat, "format", "simple", "Print detailed information in the specified format (supported formats: table, json, csv)")
	CommandDefinition.PersistentFlags().BoolVar(&configPrintSerializeWrapInObject, "json-wrap-in-object", false, "JSON: Wrap the array in an object (a compatibility option for some weak JSON parsers)")
}

func main(cmd *cobra.Command, args []string) error {
	log.Traceln("serial list main()")
	serialPortList, err := serial.GetPortsList()
	if err != nil {
		log.WithError(err).Errorln("error getting serial port list")
		return err
	}

	separator := "\n"
	if configPrint0 {
		separator = "\x00"
	}

	switch strings.ToLower(configPrintFormat) {
	case "simple":
		fmt.Print(strings.Join(serialPortList, separator))

	case "detail", "detailed", "table":
		formatTable()

	case "json":
		var d interface{}
		if configPrintSerializeWrapInObject {
			d = serialPortEnvelope{
				Ports: getPortDetail(),
			}
		} else {
			d = getPortDetail()
		}
		return portDetailDump(d, func(d interface{}) ([]byte, error) {
			return json.MarshalIndent(d, "", "    ")
		})

	case "csv":
		d := getPortDetail()
		return portDetailDump(d, csvutil.Marshal)

	default:
		log.Errorf("unknown format: %s", configPrintFormat)
		return errors.New("unknown format")
	}

	return nil
}
