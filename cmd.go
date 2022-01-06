//go:build windows

package common

import (
	"errors"
	"runtime"
	"syscall"
	"unsafe"
)

func SetCmdTitle(title string) error {
	if runtime.GOOS != "windows" {
		return errors.New("not supported os")
	}
	kernel32, _ := syscall.LoadLibrary(`kernel32.dll`)
	sct, _ := syscall.GetProcAddress(kernel32, `SetConsoleTitleW`)
	t, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return err
	}
	_, _, err = syscall.Syscall(sct, 1, uintptr(unsafe.Pointer(t)), 0, 0)
	if err != nil {
		return err
	}
	return syscall.FreeLibrary(kernel32)
}
