package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

const port = ":8080"

type worker struct {
	id       int
	alive    bool
	socket   string
	lastLoad int
	lastSeen int64
}

type pool struct {
	lastWorker int
	workers    []worker
}

type httpMessage struct {
	URL string `json:"url"`
}

var p pool

func main() {
	p.workers = []worker{}
	p.workers = append(p.workers, *NewWorker("127.0.0.1:8081"))

	http.HandleFunc("/url", handleURL)
	log.Println("Listening:", port[1:])
	log.Fatal(http.ListenAndServe(port, nil))

	p.QuitWork()
}

// handleURL - test with: curl -X POST -H 'Content-Type: application/json' -d "{\"url\": \"https://golangci.com/product#linters\"}" http://127.0.0.1:8080/url
func handleURL(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")
	decoder := json.NewDecoder(r.Body)

	var ev httpMessage
	errDecode := decoder.Decode(&ev)
	if errDecode != nil {
		log.Fatal("error decode: ", errDecode)
		return
	}
	log.Println("json: ", ev.URL)

	imgPath, errImage := requestURL(p.RequestWorker().socket, ev.URL)
	if errImage != nil {
		// HTTP error to send in header
		return
	}
	resp := httpMessage{URL: imgPath[:len(imgPath)-1]}
	jstream := json.NewEncoder(w)
	jstream.Encode(&resp)
}

func requestURL(socket, url string) (string, error) {
	conn, _ := net.Dial("tcp", socket)
	fmt.Fprintf(conn, url+"\n")

	reply, errRecv := bufio.NewReader(conn).ReadString('\n')
	log.Println("received: ", reply)
	return reply, errRecv
}

func NewWorker(sock string) *worker {
	return &worker{
		socket: sock,
	}
}
