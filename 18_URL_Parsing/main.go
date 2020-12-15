package main

import (
	"log"
	"regexp"
)

func returnID(url, regex string) (string, error) {
	r, errParse := regexp.Compile(regex)
	if errParse != nil {
		return "", errParse
	}

	vals := r.FindStringSubmatch(url)
	log.Println("Passed URL:", url, "Parsed URL:", vals)
	return vals[1], nil
}

func main() {
	url1 := "/id/123"
	x1, _ := returnID(url1, "/id/(.*)")

	log.Println("Value:", x1)

	url2 := "/id/124/abc"
	x2, _ := returnID(url2, "/id/124/(.*)")

	log.Println("Value:", x2)
}
