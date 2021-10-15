package panic_helper

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type f func()

// DoNotPanic stops a panic from escaping to the outside world.
func DoNotPanic(f f) (err error) {
	defer func() {
		recoveryErr := recover()
		if recoveryErr != nil {
			err = errors.New(fmt.Sprint(recoveryErr))
			log.WithError(err).Tracef("recovered from panic, err (%T)", recoveryErr)
		} // else no panic happened
	}()

	f()
	return nil
}
