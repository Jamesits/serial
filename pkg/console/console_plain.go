package console

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

func ConsolePlain() (in chan<- []byte, out <-chan []byte, err error) {
	var inChan = make(chan []byte)
	var outChan = make(chan []byte)

	log.Traceln("term_plain start")
	log.Warningln("using fallback terminal mode")

	go func() {
		log.Traceln("write process start")

		for buffer := range inChan {
			n, err := os.Stdout.Write(buffer)
			log.WithError(err).Tracef("term write: %d bytes", n)
			if err != nil {
				break
			}
		}

		close(inChan)
		log.Traceln("write process ended")
	}()

	go func() {
		log.Traceln("read process start")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			var buffer = scanner.Bytes()
			log.WithError(err).Tracef("term read: %d bytes", len(buffer))
			if err != nil {
				break
			}
			outChan <- []byte(buffer)
			outChan <- linebreak
		}

		close(outChan)
		log.Traceln("read process ended")
	}()

	return inChan, outChan, nil
}
