// +build darwin freebsd netbsd openbsd solaris dragonfly

package raw_term

import "syscall"

// Reference: https://github.com/cheggaaa/pb/blob/562298cc9730210a05f75d718b56c07e494965fa/termios_bsd.go
const IoctlReadTermios = syscall.TIOCGETA
const IoctlWriteTermios = syscall.TIOCSETA
