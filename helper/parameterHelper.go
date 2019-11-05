package helper

import (
	"syscall"
)

func GetParameter(name string) (string, bool) {
	return syscall.Getenv(name)
}