package main

import (
	"log"
	"testing"
)

func Test_Generator(t *testing.T) {
	n := 35
	token := customGenerator(n)
	if len(token) != 35 {
		t.Error("token generation not working")
	}
	log.Println(len(token), "--", token)
}

func Test_Random(t *testing.T) {
	token := randomNumber(10)
	if token == 0 || token > 9 {
		t.Error("random generation not working")
	}
	log.Println("generated number:", token)
}
