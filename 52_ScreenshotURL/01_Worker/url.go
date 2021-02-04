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

func createFile(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

func deleteFile(path string) error {
	return os.Remove(path)
}

func existsFile(path string) bool {
	_, errStat := os.Stat(path)
	return errStat == nil
}

func runCmd(command, args string) error {
	cmd := exec.Command(command, args)
	log.Println("running: ", cmd.Args)

	if errStart := cmd.Start(); errStart != nil {
		log.Fatal(errStart)
		return errStart
	}
	log.Println("waiting ...", cmd.Args[1])

	return cmd.Wait()
}

func hashURL(url string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(url)))[:16] // adjust length per traffic
}

func doWork(url string) (string, error) {
	urlHash := hashURL(url)
	imgPath := ramDisk + urlHash + ".png"
	scriptName := urlHash + ".sh"

	if errCreateScript := createFile(scriptName, cmd+imgPath+" "+url); errCreateScript != nil {
		return "", errors.New("script not created")
	}

	timeStart := time.Now()

	errCmd := runCmd("bash", scriptName)
	if errCmd != nil || !existsFile(imgPath) {
		return "", errors.New("screenshot not created")
	}

	elapsed := time.Since(timeStart)
	log.Println("seconds: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 1, 64))

	if errDelete := deleteFile(scriptName); errDelete != nil {
		log.Println("file: ", scriptName, " could not be deleted")
	}

	return imgPath, errCmd
}
