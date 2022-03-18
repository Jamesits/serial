package console

import (
	"github.com/jamesits/libiferr/panicked"
	log "github.com/sirupsen/logrus"
	"os"
)

func StdoutCooked(stdout *os.File) (in chan<- []byte, err error) {
	log.Debugln("stdout setup")

	var inChan = make(chan []byte)
	go func() {
		log.Debugln("write process start")
		defer panicked.Catch(func() {
			log.Debugln("closing write channel")
			close(inChan)
		})

		for buffer := range inChan {
			n, err := stdout.Write(buffer)
			log.WithError(err).Tracef("term write: %d bytes", n)
			if err != nil {
				break
			}
		}

		log.Debugln("write process ended")
	}()

	return inChan, nil
}
