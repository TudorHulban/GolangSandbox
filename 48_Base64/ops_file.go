package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func writeFile(pContent []byte, pFilePath string) error {
	return ioutil.WriteFile(pFilePath, pContent, 0644)
}

func calcMD5(pFilePath string) (string, error) {
	fileHandler, errOpenFile := os.Open(pFilePath)
	if errOpenFile != nil {
		return "", errOpenFile
	}
	defer fileHandler.Close()

	stat, _ := fileHandler.Stat()
	log.Println("md5 for:", pFilePath, "size:", stat.Size(), "name:", stat.Name())

	hash := md5.New()
	_, errCopy := io.Copy(hash, fileHandler)
	if errCopy != nil {
		return "", errCopy
	}

	var result string
	resultBytes := hash.Sum(nil)[:16]
	result = hex.EncodeToString(resultBytes)
	return result, nil
}
