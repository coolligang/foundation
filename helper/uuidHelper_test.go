package helper

import "testing"

func TestGetUUID(t *testing.T) {
	str, _ := GetUUID(16)
	t.Log(str)
	if len(str) != 2*16 {
		t.Fail()
	}
}
