package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func createFile(path string, fileContent string) error {
	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func existsFile(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// runCmd Helper to bash execute a commsn with given arguments.
// introduced as syscall did not work with firefox.
func runCmd(command, theArguments string) error {
	cmd := exec.Command(command, theArguments)

	log.Println("running: ", cmd.Args)
	if errStart := cmd.Start(); errStart != nil {
		log.Fatal(errStart)
	}

	log.Println("waiting ...", cmd.Args[1])
	return cmd.Wait()
}

// hashURL Helper to sha the URL in order to have a unique name for the screen capture.
func hashURL(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:16] // adjust length per needs
}
