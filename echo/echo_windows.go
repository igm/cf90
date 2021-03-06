// +build windows

package echo

import (
	"fmt"
	"os"
	"syscall"
)

// see SetConsoleMode documentation for bit flags
// http://msdn.microsoft.com/en-us/library/windows/desktop/ms686033(v=vs.85).aspx
const ENABLE_ECHO_INPUT = 0x0004

func echoOff(f func()) {
	hStdin := syscall.Handle(os.Stdin.Fd())
	var originalMode uint32

	err := syscall.GetConsoleMode(hStdin, &originalMode)
	if err != nil {
		return
	}
	var newMode uint32 = (originalMode &^ ENABLE_ECHO_INPUT)

	err = setConsoleMode(hStdin, newMode)
	if err != nil {
		return
	}
	f()
	setConsoleMode(hStdin, originalMode)
	fmt.Println()
}

func setConsoleMode(console syscall.Handle, mode uint32) (err error) {
	dll := syscall.MustLoadDLL("kernel32")
	proc := dll.MustFindProc("SetConsoleMode")
	r, _, err := proc.Call(uintptr(console), uintptr(mode))

	if r == 0 {
		return err
	}
	return nil
}
