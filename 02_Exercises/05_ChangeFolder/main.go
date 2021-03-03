package main

import (
	"strings"

	"fmt"
)

var currentFolder string

func main() {
	currentFolder = "/"

	input := make([]string, 0)
	input = append(input, "pwd")
	input = append(input, "cd /home/csed")
	input = append(input, "pwd")
	input = append(input, "cd /lab/root/../dir")
	input = append(input, "pwd")
	input = append(input, "cd /lab/root/../dir/../../x")
	input = append(input, "pwd")
	input = append(input, "cd /home")
	input = append(input, "pwd")
	input = append(input, "cd lab")
	input = append(input, "pwd")

	process(input)

}

func process(input []string) {
	for _, v := range input {

		if v == "pwd" {
			fmt.Println("pwd:", currentFolder)
			continue
		}

		nextPath := strings.Split(v, " ")[1]

		currentFolder = handlePathLogic(nextPath)
	}
}

func handlePathLogic(p string) string {
	if len(p) == 0 {
		return ""
	}

	if p[0] != '/' {
		return currentFolder + "/" + p
	}

	elems := strings.Split(strings.Trim(p, " "), "/")

	var result []string

	for _, v := range elems {
		if len(v) == 0 {
			continue
		}

		if v == ".." {
			result = result[:len(result)-1]

			if len(result) == 0 {
				result = []string{"/"}
			}

			if len(strings.Trim(result[0], " ")) == 0 {
				result = []string{"/"}
			}

			continue
		}

		result = append(result, v)
	}

	if result[0] == "/" {
		return "/" + strings.Join(result[1:], "/")
	}

	return "/" + strings.Join(result, "/")
}
