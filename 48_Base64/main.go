/*
encode, decode file. store as encoded in postgres as text column.
tested with up to 7.5 MB.
no issues detected.
*/
package main

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const sourcePath = "assets/k3.pdf"
const targetPath = "processed/k3.pdf"
const dbID = 2

func main() {
	f, _ := os.Open(sourcePath)
	defer f.Close()
	content, _ := ioutil.ReadAll(bufio.NewReader(f))
	encoded := base64.StdEncoding.EncodeToString(content)

	m1, err1 := calcMD5(sourcePath)
	if err1 != nil {
		log.Println("err1:", err1)
		os.Exit(30)
	}
	var file File
	file.ID = dbID
	file.Name = filepath.Base(sourcePath)
	file.Content = encoded
	insertFile(&file)

	fileDB, _ := fetchFile(dbID)
	contentDB := fileDB.Content
	decoded, _ := base64.StdEncoding.DecodeString(contentDB)

	writeFile(decoded, targetPath)

	m2, err2 := calcMD5(targetPath)
	if err2 != nil {
		log.Println("err2:", err2)
		os.Exit(47)
	}
	log.Println("md5: m1", m1, err1)
	log.Println("md5: m2", m2, err2)
	log.Println("files matching:", m1 == m2)
}
