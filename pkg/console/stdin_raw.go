package console

import (
	"github.com/Jamesits/serial/pkg/raw_term"
	"github.com/jamesits/libiferr/panicked"
	log "github.com/sirupsen/logrus"
)

// https://viewsourcecode.org/snaptoken/kilo/03.rawInputAndOutput.html
func StdinRaw() (out <-chan []byte, err error) {
	log.Debugln("stdin_raw setup")
	err, _ = panicked.CatchError(func() {
		raw_term.SetRaw()
	})
	if err != nil {
		return
	}

	var outChan = make(chan []byte)
	go func() {
		log.Debugln("read process start")
		defer raw_term.Restore()
		defer panicked.Catch(func() {
			log.Debugln("closing read channel")
			close(outChan)
		})

		for true {
			buffer, err := raw_term.ReadByteSync()
			log.WithError(err).WithField("len", len(buffer)).WithField("content", buffer).Tracef("term read")
			if err != nil || len(buffer) != 1 {
				break
			}

			// TODO: make a big keyboard state machine here for shortcut & escape key processing
			if buffer[0] == CtrlKey('c') {
				log.Traceln("^C detected")
				break
			}

			outChan <- buffer
		}

		log.Debugln("read process ended")
	}()

	return outChan, nil
}

// CtrlKey makes a Ctrl-* combination from the input ASCII character.
func CtrlKey(key byte) byte {
	return key & 0x1f
}
