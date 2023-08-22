package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash(val string) string {
	hash := sha256.New().Sum([]byte(val))
	return fmt.Sprintf("%x", hash)
}
