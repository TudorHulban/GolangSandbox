package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const thePath = "/home/tudi/ram"

func main() {
	allFiles := make([]string, 0)

	doPathWalk(thePath, &allFiles) //AllFiles now contains files in subfolders of APath

	for _, f := range allFiles {
		fmt.Println(f)
	}
}

func doPathWalk(folder string, files *[]string) error {
	fullPath, err := filepath.Abs(folder)
	if err != nil {
		return err
	}

	callback := func(path string, fi os.FileInfo) error {
		return hashFile(fullPath, path, fi, files)
	}

	return filepath.Walk(fullPath, callback)
}

func hashFile(root, path string, fi os.FileInfo, files *[]string) error {
	if fi.IsDir() {
		return nil
	}

	if _, err := filepath.Rel(root, path); err != nil {
		return err
	}

	*files = append(*files, path)

	return nil
}
