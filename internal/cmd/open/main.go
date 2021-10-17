package open

import (
	"github.com/Jamesits/serial/pkg/console"
	"github.com/Jamesits/serial/pkg/panic_helper"
	"github.com/mattn/go-isatty"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.bug.st/serial"
	"os"
	"strings"
	"sync"
)

var CommandDefinition = &cobra.Command{
	Use:   "open",
	Short: "Opens a serial port",
	Args:  cobra.ExactArgs(1),
	RunE:  main,
}

var serialMode serial.Mode
var configParityMode string
var configStopBits string
var mainThreadWaitGroup = &sync.WaitGroup{}

func init() {
	CommandDefinition.PersistentFlags().IntVarP(&serialMode.BaudRate, "baudrate", "b", 9600, "Symbol rate of the port")
	CommandDefinition.PersistentFlags().IntVarP(&serialMode.DataBits, "data-bits", "d", 8, "Size of the character")
	CommandDefinition.PersistentFlags().StringVarP(&configParityMode, "parity", "p", "no", "Parity mode (no, odd, even, mark, space)")
	CommandDefinition.PersistentFlags().StringVarP(&configStopBits, "stop-bits", "s", "1", "Stop bits (1, 1,5, 2)")
}

func main(cmd *cobra.Command, args []string) error {
	var err error
	log.Traceln("verifying config")

	// baudrate
	if serialMode.BaudRate <= 0 {
		log.Errorf("unsupported baudrate: %d", serialMode.BaudRate)
		return ErrorUnknownConfig
	}

	// data bits
	if serialMode.DataBits < 5 || serialMode.DataBits > 8 {
		log.Errorf("unsupported data bits: %d", serialMode.DataBits)
		return ErrorUnknownConfig
	}

	// parity
	switch strings.ToLower(strings.TrimSpace(configParityMode)) {
	case "no":
	case "n":
		serialMode.Parity = serial.NoParity

	case "odd":
	case "o":
		serialMode.Parity = serial.OddParity

	case "even":
	case "e":
		serialMode.Parity = serial.EvenParity

	case "mark":
	case "m":
		serialMode.Parity = serial.MarkParity

	case "space":
	case "s":
		serialMode.Parity = serial.SpaceParity

	default:
		log.Errorf("unknown parity mode: %s", configParityMode)
		return ErrorUnknownConfig
	}

	// stop bits
	switch strings.ToLower(strings.TrimSpace(configStopBits)) {
	case "1":
		serialMode.StopBits = serial.OneStopBit

	case "1.5":
		serialMode.StopBits = serial.OnePointFiveStopBits

	case "2":
		serialMode.StopBits = serial.TwoStopBits

	default:
		log.Errorf("unknown stop bits: %s", configStopBits)
		return ErrorUnknownConfig
	}

	// setup I/O
	var in chan<- []byte
	var out <-chan []byte
	err = nil

	var stdinIsATerminal = true
	if isatty.IsTerminal(os.Stdin.Fd()) {
		log.Debugln("stdin is a terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		log.Debugln("stdin is a terminal")
	} else {
		stdinIsATerminal = false
		log.Debugln("stdin is not a terminal")
	}

	if stdinIsATerminal {
		out, err = console.StdinRaw()
	} else {
		// ???
	}

	in, err = console.StdoutCooked(os.Stdout)

	//for _, console := range console.AvailableConsoleTypes {
	//	in, out, err = console()
	//	if err == nil {
	//		break
	//	} else {
	//		log.WithError(err).Debugln("unable to initialize console type %v", console)
	//	}
	//}
	//if err != nil {
	//	log.Errorln("no console options available")
	//	return err
	//}

	log.Tracef("opening serial port")
	sp, err := serial.Open(args[0], &serialMode)
	if err != nil {
		log.WithError(err).Errorln("open serial port failed")
		return err
	}
	log.Info("serial port opened")

	// read
	mainThreadWaitGroup.Add(1)
	go func() {
		var buffer = make([]byte, 256)
		var n int
		var err error

		log.Traceln("serial read process start")

		for true {
			n, err = sp.Read(buffer)
			log.WithError(err).Tracef("serial read: %d bytes", n)
			if err != nil {
				break
			}

			in <- buffer[:n]
		}

		log.Traceln("serial read process ended")
		panic_helper.DoNotPanic(func() {
			close(in)
		})
		mainThreadWaitGroup.Done()
	}()

	// write
	mainThreadWaitGroup.Add(1)
	go func() {
		var n int
		var err error

		log.Traceln("serial write process start")

		for buffer := range out {
			n, err = sp.Write(buffer)
			log.WithError(err).Tracef("serial write: %d bytes", n)
			if err != nil {
				break
			}
		}

		log.Traceln("serial write process ended")

		// signal serial read process to quit
		panic_helper.DoNotPanic(func() {
			close(in)
		})

		mainThreadWaitGroup.Done()
	}()

	mainThreadWaitGroup.Wait()
	err = sp.Close()
	if err != nil {
		log.WithError(err).Errorln("error closing serial port")
	}
	return nil
}
