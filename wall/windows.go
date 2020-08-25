// +build windows

package wall

import (
	"syscall"
)

func GetWinScreenSize(nIndex int) (i int) {
	rs, _, _ := syscall.NewLazyDLL("User32.dll").NewProc("GetSystemMetrics").Call(uintptr(nIndex))
	i = int(rs)
	return i
}