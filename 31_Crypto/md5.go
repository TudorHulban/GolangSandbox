package main

import (
	"crypto/md5"
	"encoding/hex"
)

// createMD5 - Returns the MD5 of passed string.
func createMD5(s string) string {
	m := md5.New()
	m.Write([]byte(s))

	return hex.EncodeToString(m.Sum(nil))
}
