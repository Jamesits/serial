package logging

import (
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

type preQuitHandler func()

func InitGlobalLogger(preQuitHandler preQuitHandler) {
	// Initialize logging
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})

	// log.SetOutput(os.Stderr) // works for *nix only
	log.SetOutput(colorable.NewColorableStderr()) // adds Windows conhost.exe compatibility

	log.SetLevel(log.InfoLevel)
	//log.SetLevel(log.TraceLevel)

	if preQuitHandler != nil {
		log.RegisterExitHandler(preQuitHandler)
		initSigintSupport(preQuitHandler)
	}
}

var SigIntChannel = make(chan os.Signal, 1)

func initSigintSupport(preQuitHandler preQuitHandler) {
	if preQuitHandler == nil {
		return
	}

	// handle SIGINT
	log.Traceln("init SIGINT handler")
	signal.Notify(SigIntChannel, os.Interrupt)
	go func() {
		for range SigIntChannel {
			preQuitHandler()
			os.Exit(0)
		}
	}()
}
