package main

import (
	"log"
	"testing"
)

func TestHashing(t *testing.T) {
	x1 := hashURL("www.olx.ro")
	x2 := hashURL("www.olx.ro/")

	log.Println(x1, "-------", len(x1))
	log.Println(x2, "-------", len(x2))
}

func TestFileExists(t *testing.T) {
	name := "xxx.png"

	errCreate := createFile(name, "xxx")
	if errCreate != nil || !existsFile(name) {
		t.Error("failed file exists")
	}
}
