package panic_helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoNotPanic(t *testing.T) {
	defer func() {
		err := recover()
		assert.Nil(t, err)
	}()

	DoNotPanic(func() {
		panic("try panic")
	})
}
