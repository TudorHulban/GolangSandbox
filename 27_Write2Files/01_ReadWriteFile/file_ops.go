package main

import (
	"bufio"
	"log"
	"os"
)

// ReadFile Helper reading file contents to a slice.
func ReadFile(filePath string) ([]string, error) {
	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileHandler.Close()

	var result = make([]string, 0)

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
}

// LogToFile Helper inserts log message in file and closes.
func LogToFile(filePath, text, prefix string) error {
	fileHandler, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fileHandler.Close()

	logger := log.New(fileHandler, prefix, log.LstdFlags)
	logger.Println(text)
	return nil
}
