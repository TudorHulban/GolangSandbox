package main

import (
	"golang.org/x/crypto/bcrypt"
)

func Create32HASH(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), 14)
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
