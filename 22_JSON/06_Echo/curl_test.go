package main

import (
	"log"
	"os/exec"
	"testing"
)

// TestSendJSON - does not work
func TestSendJSON(t *testing.T) {
	params := prepareCurl("-X POST", "-H 'Content-Type: application/json'", "-d \"{\\\"message\\\": \\\"Hi\\\"}\"", "http://127.0.0.1:8080/url")
	ctx := exec.Command("curl", params)
	log.Println("ctx: ", ctx.Args)

	errCurl := ctx.Run()
	if errCurl != nil {
		t.Error("curl error: ", errCurl)
	}
}

func prepareCurl(pParameters ...string) string {
	var result string
	for _, v := range pParameters {
		result = result + " " + v
	}
	return result
}
