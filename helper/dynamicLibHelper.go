package helper

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetDynamicLibrary(path, nameWithoutSuffix string) (ptr uintptr) {
	if runtime.GOOS == "windows" {
		fmt.Println(filepath.Join(path, nameWithoutSuffix+".dll"))
		ptr = LoadDynamicLibrary(filepath.Join(path, nameWithoutSuffix+".dll"))
	} else {
		fmt.Println(filepath.Join(path, nameWithoutSuffix+".so"))
		ptr = LoadDynamicLibrary(filepath.Join(path, nameWithoutSuffix+".so"))
	}
	return ptr
}

func GetFunctionInDynamicLibrary(lib uintptr, name string) uintptr {
	if lib != 0 {
		return GetFuncAddrInLibrary(lib, name)
	}
	return 0
}
