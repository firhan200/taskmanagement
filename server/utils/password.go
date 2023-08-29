package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashPassword(p string) (string, error) {
	sha1 := sha1.New()
	sha1.Write([]byte(p))
	sha1_hash := hex.EncodeToString(sha1.Sum(nil))
	return sha1_hash, nil
}
