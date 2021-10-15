//go:build darwin || linux
// +build darwin linux

// https://gist.github.com/EddieIvan01/4449b64fc1eb597ffc2f317cfa7cc70c

package raw_term

import (
	log "github.com/sirupsen/logrus"
	"os"
	"syscall"
	"unsafe"
)

func getTermios(fd uintptr) *syscall.Termios {
	var t syscall.Termios
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCGETS,
		uintptr(unsafe.Pointer(&t)),
		0, 0, 0)

	if err != 0 {
		panic("err")
	}

	return &t
}

func setTermios(fd uintptr, term *syscall.Termios) {
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,
		os.Stdin.Fd(),
		syscall.TCSETS,
		uintptr(unsafe.Pointer(term)),
		0, 0, 0)
	if err != 0 {
		panic("err")
	}
}

func setRaw(term *syscall.Termios) {
	// This attempts to replicate the behaviour documented for cfmakeraw in
	// the termios(3) manpage.
	term.Iflag &^= syscall.IGNBRK | syscall.BRKINT | syscall.PARMRK | syscall.ISTRIP | syscall.INLCR | syscall.IGNCR | syscall.ICRNL | syscall.IXON
	// newState.Oflag &^= syscall.OPOST
	term.Lflag &^= syscall.ECHO | syscall.ECHONL | syscall.ICANON | syscall.ISIG | syscall.IEXTEN
	term.Cflag &^= syscall.CSIZE | syscall.PARENB
	term.Cflag |= syscall.CS8

	term.Cc[syscall.VMIN] = 1
	term.Cc[syscall.VTIME] = 0
}

var origin syscall.Termios

func SetRaw() {
	t := getTermios(os.Stdin.Fd())
	origin = &t

	setRaw(t)
	setTermios(os.Stdin.Fd(), t)
}

func Restore() {
	setTermios(os.Stdin.Fd(), &origin)
}

func ReadByteSync() ([]byte, error) {
	buf := make([]byte, 1)
	n, err := syscall.Read(0, buf)
	log.WithError(err).WithField("len", n).Traceln("raw read from stdin")

	// identify pipe close
	if err == nil && n == 0 {
		err = ErrorPipeBroken
	}
	return buf, err
}
