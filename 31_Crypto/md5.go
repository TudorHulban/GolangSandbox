package main

import (
	"crypto/md5"
	"encoding/hex"
)

func createMD5(pString string) string {

	m := md5.New()
	m.Write([]byte(pString))
	return hex.EncodeToString(m.Sum(nil))
}
