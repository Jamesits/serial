// +build windows

package raw_term

import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32.dll")

var (
	procGetConsoleMode = kernel32.NewProc("GetConsoleMode")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
)

const (
	enableLineInput       = 2
	enableEchoInput       = 4
	enableProcessedInput  = 1
	enableWindowInput     = 8
	enableMouseInput      = 16
	enableInsertMode      = 32
	enableQuickEditMode   = 64
	enableExtendedFlags   = 128
	enableAutoPosition    = 256
	enableProcessedOutput = 1
	enableWrapAtEolOutput = 2
)

func getTermMode(fd uintptr) uint32 {
	var mode uint32
	_, _, err := syscall.Syscall(
		procGetConsoleMode.Addr(),
		2,
		fd,
		uintptr(unsafe.Pointer(&mode)),
		0)
	if err != 0 {
		panic("err")
	}

	return mode
}

func setTermMode(fd uintptr, mode uint32) {
	_, _, err := syscall.Syscall(
		procSetConsoleMode.Addr(),
		2,
		fd,
		uintptr(mode),
		0)
	if err != 0 {
		panic("err")
	}
}

var origin uint32

func SetRaw() {
	originMode := getTermMode(os.Stdin.Fd())
	origin = originMode

	originMode &^= enableEchoInput | enableProcessedInput | enableLineInput | enableProcessedOutput
	setTermMode(os.Stdin.Fd(), originMode)
}

func Restore() {
	setTermMode(os.Stdin.Fd(), origin)
}

func ReadByteSync() ([]byte, error) {
	buf := make([]byte, 1)
	n, err := syscall.Read(syscall.Handle(os.Stdin.Fd()), buf)
	log.WithError(err).WithField("len", n).Traceln("raw read from stdin")

	// identify pipe close
	if err == nil && n == 0 {
		err = ErrorPipeBroken
	}
	return buf, err
}
