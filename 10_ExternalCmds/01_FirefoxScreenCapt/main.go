package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

const ramDisk = ""
const cmd = "firefox -headless --screenshot "

const (
	u1 = "https://www.olx.ro/oferta/lifepo4-acumulator-baterie-incarcator-60ah-100ah-200ah-ID9XZiR.html"
	u2 = "https://www.olx.ro/oferta/vand-stupi-albine-IDctzni.html"
)

func init() {
	runCmd("bash", "cleanpng.sh")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	doWork(&wg, u1)
	doWork(&wg, u2)

	wg.Wait()
}

func doWork(wg *sync.WaitGroup, urltoCapture string) {
	defer wg.Done()

	urlHash := hashURL(urltoCapture)
	imgPath := ramDisk + urlHash + ".png"

	scriptName := urlHash + ".sh"
	createFile(scriptName, cmd+imgPath+" "+urltoCapture)

	t := time.Now()
	runCmd("bash", scriptName)

	elapsed := time.Since(t)
	log.Println("seconds: " + strconv.FormatFloat(elapsed.Seconds(), 'f', 1, 64))

	log.Println(imgPath + " exists ?: " + strconv.FormatBool(existsFile(imgPath)))
	deleteFile(scriptName)
}
