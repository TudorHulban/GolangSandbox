package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func createFile(filePath string, content string) error {
	// TODO: verify if file exists, maybe add override param.
	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

func deleteFile(filePath string) error {
	if fileExists(filePath) == false {
		return nil // file does not exist, nothing to do.
	}

	return os.Remove(filePath)
}

// fileExists Helper for checking if we receive specific error for file path.
// https://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
func fileExists(filePath string) bool {
	if _, errPath := os.Stat(filePath); errPath != nil {
		// check received error
		if os.IsNotExist(errPath) {
			return false
		}
	}

	return true
}

func runCmd(command, arguments string) error {
	cmd := exec.Command(command, arguments)
	log.Println("running: ", cmd.Args)

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	log.Println("waiting ...", cmd.Args[1])
	return cmd.Wait()
}

func hashURL(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:16] // adjust length per needs
}

func doWork(url string) (string, error) {
	urlHash := hashURL(url)
	imgPath := ramDisk + urlHash + ".png"
	scriptName := urlHash + ".sh"
	createFile(scriptName, cmd+imgPath+" "+url)

	t1 := time.Now()
	errCmd := runCmd("bash", scriptName)
	t2 := time.Now()

	elapsed := t2.Sub(t1)
	if fileExists(imgPath) == false {
		return "", errors.New("file not created")
	}

	log.Println("seconds: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 1, 64))
	log.Println(imgPath + " exists ?: " + strconv.FormatBool(fileExists(imgPath)))
	deleteFile(scriptName)

	return imgPath, errCmd
}
