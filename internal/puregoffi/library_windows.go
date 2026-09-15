//go:build windows

package puregoffi

import "syscall"

func OpenLibrary(name string) (uintptr, error) {
	handle, err := syscall.LoadLibrary(name)
	return uintptr(handle), err
}

func OpenSymbol(library uintptr, name string) (uintptr, error) {
	return syscall.GetProcAddress(syscall.Handle(library), name)
}

func CloseLibrary(library uintptr) error {
	return syscall.FreeLibrary(syscall.Handle(library))
}
