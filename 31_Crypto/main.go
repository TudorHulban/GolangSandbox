package main

import (
	"fmt"
)

func main() {
	password := "secret"
	key := []byte(createMD5(password))

	aes, _ := EncryptAES(key, "xxxxxxxxxxxx")
	msg, _ := DecryptAES(key, aes)

	fmt.Println("message:", msg)
}
