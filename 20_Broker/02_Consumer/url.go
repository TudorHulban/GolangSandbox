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

func createFile(pPath string, pContent string) error {
	content := []byte(pContent)
	return ioutil.WriteFile(pPath, content, 0644)
}

func deleteFile(pPath string) error {
	return os.Remove(pPath)
}

func existsFile(pPath string) bool {
	_, errStat := os.Stat(pPath)
	return errStat == nil
}

func runCmd(pCommand, pArgs string) error {
	cmd := exec.Command(pCommand, pArgs)
	log.Println("running: ", cmd.Args)

	errStart := cmd.Start()
	if errStart != nil {
		log.Fatal(errStart)
		return errStart
	}
	log.Println("waiting ...", cmd.Args[1])
	errWait := cmd.Wait()
	return errWait
}

func hashURL(pURL string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(pURL)))[:16] // adjust length per traffic
}

func doWork(pURL string) (string, error) {
	urlHash := hashURL(pURL)
	imgPath := ramDisk + urlHash + ".png"
	scriptName := urlHash + ".sh"
	errCreateScript := createFile(scriptName, cmd+imgPath+" "+pURL)
	if errCreateScript != nil {
		return "", errors.New("script not created")
	}

	timeStart := time.Now()
	errCmd := runCmd("bash", scriptName)
	if errCmd != nil || !existsFile(imgPath) {
		return "", errors.New("screenshot not created")
	}
	elapsed := time.Since(timeStart)
	log.Println("seconds: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 1, 64))

	errDelete := deleteFile(scriptName)
	if errDelete != nil {
		log.Println("file: ", scriptName, " could not be deleted")
	}
	return imgPath, errCmd
}
