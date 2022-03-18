package console

import (
	"errors"
	"github.com/containerd/console"
	"github.com/jamesits/libiferr/panicked"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

// https://stackoverflow.com/a/58237997
func ConsoleRawExperimental() (in chan<- []byte, out <-chan []byte, err error) {
	log.Tracef("term_raw_experimental setup console")
	defer func() {
		ret := recover()
		if ret != nil {
			log.Warningf("set console raw mode failed: %s", err)
			err = errors.New("unable to set console raw mode")
		}
	}()

	current := console.Current()
	if err = current.SetRaw(); err != nil {
		log.WithError(err).Debugln("console set raw mode failed")
		return
	}

	// disable echo
	// TODO: write a terminal emulator and detect if echo need to be enabled
	if err = current.DisableEcho(); err != nil {
		log.WithError(err).Debugln("console disable echo failed")
		return
	}

	term := terminal.NewTerminal(current, "")
	term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
		log.WithField("line", line).WithField("pos", pos).WithField("key", key).Traceln("term.AutoCompleteCallback")
		return "", 0, false
	}

	var inChan = make(chan []byte)
	var outChan = make(chan []byte)

	go func() {
		log.Traceln("write process start")

		for buffer := range inChan {
			n, err := term.Write(buffer)
			log.WithError(err).Tracef("term write: %d bytes", n)
			if err != nil {
				break
			}
		}

		panicked.Catch(func() {
			close(inChan)
		})
		log.Traceln("write process ended")
	}()

	go func() {
		log.Traceln("read process start")
		defer func(current console.Console) {
			err := current.Reset()
			if err != nil {
				log.WithError(err).Warningln("console reset failed")
			}
		}(current)

		for true {
			line, err := term.ReadLine()
			log.WithError(err).Tracef("term read: %d bytes", len(line))
			if err != nil {
				break
			}
			outChan <- []byte(line)
			outChan <- linebreak
		}

		panicked.Catch(func() {
			close(outChan)
		})

		// FIXME: correctly signal inChan to close
		panicked.Catch(func() {
			close(inChan)
		})

		log.Traceln("read process ended")
	}()

	return inChan, outChan, nil
}
