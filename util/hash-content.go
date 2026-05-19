// Package util conatins helper files
package util

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashContent(data []byte) string {
	sum := sha1.Sum(data)                 // Array [20]byte
	encoded := hex.EncodeToString(sum[:]) // [:] gjør om array til []byte
	return encoded
}
