package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	waitSecs = 4
	slowURL  = "https://httpstat.us/200"
	fastURL  = "https://olx.ro"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(waitSecs)*time.Second)
	defer cancel()

	request, errReq := http.NewRequest(http.MethodGet, slowURL, nil)
	if errReq != nil {
		log.Fatal(errReq)
	}

	log.Println("performing request to", slowURL)
	response, errResp := http.DefaultClient.Do(request.WithContext(ctx))
	if errResp != nil {
		log.Fatal(errResp)
	}

	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body)
}
