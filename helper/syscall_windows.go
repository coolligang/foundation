package helper

import (
	"syscall"
)

func LoadDynamicLibrary(name string) uintptr {
	defer HandleErr()
	lib, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}
	return uintptr(lib)
}

func GetFuncAddrInLibrary(lib uintptr, name string) uintptr {
	defer HandleErr()
	addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
	if err != nil {
		panic(err)
	}
	return uintptr(addr)
}
