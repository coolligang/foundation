package haspkey

import (
	"lhr/foundation/definitions"
	"os"
	"path/filepath"
	"testing"
)

func TestHaspKeyValidity(t *testing.T) {
	_ = os.Setenv(definitions.DLhrKeyHaspLibPath, filepath.Join(os.Getenv("GOPATH"), "src", "lhr/foundation/key/haspkey/api"))
	res, _ := IsCwValidity()
	t.Log(res)
	_ = os.Unsetenv(definitions.DLhrKeyHaspLibPath)
}

func TestHaspKeySupportProduct(t *testing.T) {
	_ = os.Setenv(definitions.DLhrKeyHaspLibPath, filepath.Join(os.Getenv("GOPATH"), "src", "lhr/foundation/key/haspkey/api"))
	res, _ := IsCwProduct(1)
	t.Log(res)
	_ = os.Unsetenv(definitions.DLhrKeyHaspLibPath)
}
