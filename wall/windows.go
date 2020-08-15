// +build windows

package wall

import (
	"syscall"
)

func GetWinScreenSize(nIndex int) (i int) {
	//defer func() {
	//	log.Print("无法获取屏幕分辨率")
	//}()

	rs, _, _ := syscall.NewLazyDLL("User32.dll").NewProc("GetSystemMetrics").Call(uintptr(nIndex))
	i = int(rs)
	return i
}