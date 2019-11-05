package helper

import (
	"testing"
)

func TestFileInStringChanger(t *testing.T) {
	in := `{"img": @{testdata/test.jpg}}`
	_, err := FileInStringChanger(in)
	if err != nil {
		t.Fail()
	}
}
