package helper

import (
	"crypto/rand"
	"fmt"
)

func GetUUID(len int) (string, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	uuid := fmt.Sprintf("%x", b)
	return uuid, err
}
