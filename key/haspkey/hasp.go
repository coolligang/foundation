package haspkey

/*
#cgo CFLAGS:  -I./api

#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h>

#include"authorize_product.h"

typedef int (*fn_cwIsProduct)(unsigned char ucProduct);
typedef int (*fn_cwGetValidity)(int iType);

int c_cwIsProduct(void* fn, unsigned char ucProduct)
{
	return ((fn_cwIsProduct)fn)(ucProduct);
}

int c_cwGetValidity(void* fn)
{
	return ((fn_cwGetValidity)fn)(1);
}

*/
import "C"

import (
	"fmt"
	"lhr/foundation/definitions"
	"lhr/foundation/helper"
	"unsafe"
)

var (
	libHaspSdk    uintptr
	cwIsProduct   uintptr
	cwGetValidity uintptr
)

func libHaspSdkInit() {
	libPath, _ := helper.GetParameter(definitions.DLhrKeyHaspLibPath)
	if len(libPath) == 0 {
		libPath = helper.GetCurrentPath() + "/lib"
	}
	libHaspSdk = helper.GetDynamicLibrary(libPath, "libauthorize_product_x64")
	if libHaspSdk != 0 {
		cwIsProduct = helper.GetFunctionInDynamicLibrary(libHaspSdk, "cwIsProduct")
		cwGetValidity = helper.GetFunctionInDynamicLibrary(libHaspSdk, "cwGetValidity")
	}
}

func IsCwProduct(ucProduct uint8) (C.int, error) {
	if libHaspSdk == 0 {
		libHaspSdkInit()
	}
	if libHaspSdk != 0 {
		res := C.c_cwIsProduct(unsafe.Pointer(cwIsProduct), C.uchar(ucProduct))
		return res, nil
	}
	return 0, fmt.Errorf("hasp library load failed")
}

func IsCwValidity() (C.int, error) {
	if libHaspSdk == 0 {
		libHaspSdkInit()
	}
	if libHaspSdk != 0 {
		res := C.c_cwGetValidity(unsafe.Pointer(cwGetValidity))
		return res, nil
	}
	return 0, fmt.Errorf("hasp library load failed")
}
