package main

import (
	"golang.org/x/crypto/bcrypt"
)

func Create32HASH(pString string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(pString), 14)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
