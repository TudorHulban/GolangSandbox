package archiving

/*
File implements zipping. Should go in own package once multiple implementations.
*/

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/TudorHulban/log"
)

const zipExtension = ".zip"

// Zip Struct define for building a zip operations package.
type Zip struct {
	l *log.LogInfo
}

// NewZip Constructor for package.
func NewZip(logger *log.LogInfo) IArchive {
	return &Zip{
		l: logger,
	}
}

// CompressFile Method compressing a file using zip method.
func (z *Zip) CompressFile(filePath string) error {
	// TODO: add logic
	return nil
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

// createFile Helper that returns an error if file not created.
// If created it returns a file handler and a closer function.
func createFile(filePath string) (*os.File, func(), error) {
	targetFile, err := os.Create(filePath + zipExtension)
	if err != nil {
		return nil, nil, err
	}
	defer targetFile.Close()

	// TODO: add logic
	return nil, nil, err
}

func ZipFile(filePath string) error {
	zipWriter := zip.NewWriter(targetFile)
	defer zipWriter.Close()

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filePath)
	header.Method = zip.Deflate

	fileWriter, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(fileWriter, file)
	return err
}
