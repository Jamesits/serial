// +build linux

package raw_term

import "syscall"

// Reference: https://github.com/cheggaaa/pb/blob/562298cc9730210a05f75d718b56c07e494965fa/termios_linux.go
const IoctlReadTermios = syscall.TCGETS
const IoctlWriteTermios = syscall.TCSETS
