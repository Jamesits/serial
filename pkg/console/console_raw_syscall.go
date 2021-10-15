package console

import (
	"github.com/Jamesits/serial/pkg/panic_helper"
	"github.com/Jamesits/serial/pkg/raw_term"
	log "github.com/sirupsen/logrus"
	"os"
)

func ConsoleRawSyscall() (in chan<- []byte, out <-chan []byte, err error) {
	var inChan = make(chan []byte)
	var outChan = make(chan []byte)

	log.Traceln("term_raw_syscall setup")
	err = panic_helper.DoNotPanic(func() {
		raw_term.SetRaw()
	})
	if err != nil {
		return
	}

	go func() {
		log.Traceln("write process start")
		defer panic_helper.DoNotPanic(func() {
			close(inChan)
		})

		for buffer := range inChan {
			n, err := os.Stdout.Write(buffer)
			log.WithError(err).Tracef("term write: %d bytes", n)
			if err != nil {
				break
			}
		}

		log.Traceln("write process ended")
	}()

	go func() {
		log.Traceln("read process start")
		defer raw_term.Restore()
		defer panic_helper.DoNotPanic(func() {
			close(outChan)
		})

		for true {
			buffer, err := raw_term.ReadByteSync()
			log.WithError(err).Tracef("term read: %d bytes", len(buffer))
			if err != nil {
				break
			}
			outChan <- buffer
		}

		log.Traceln("read process ended")
	}()

	return inChan, outChan, nil
}
