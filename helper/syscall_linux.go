package helper

import (
	"fmt"
	"unsafe"
)

/*
#cgo CFLAGS:  -I.
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
*/
import "C"

func LoadDynamicLibrary(name string) uintptr {
	defer HandleErr()
	handle := C.dlopen(C.CString(name), 1)
	if handle == unsafe.Pointer(nil) {
		panic(fmt.Sprintf("dlopen[%s] failed", name))
	}
	return uintptr(handle)
}

func GetFuncAddrInLibrary(lib uintptr, name string) uintptr {
	defer HandleErr()
	addr := C.dlsym(unsafe.Pointer(lib), C.CString(name))
	if addr == unsafe.Pointer(nil) {
		panic(fmt.Sprintf("dlsym[%s] failed", name))
	}
	return uintptr(addr)
}
