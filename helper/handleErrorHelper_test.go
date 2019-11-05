package helper

import (
	"testing"
	"time"
)

func TestUnitHandleErr(t *testing.T) {
	go func() {
		defer HandleErr()

		v := 0
		v = 1 / v
	}()
	time.Sleep(time.Second)

	err := recover()
	if err != nil {
		t.Fail()
	}
}
