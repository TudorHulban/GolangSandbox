package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const port = ":8080"

type httpMessage struct {
	Msg string `json:"message"`
}

func main() {
	http.HandleFunc("/url1", responseText)
	http.HandleFunc("/url2", responseJSON)
	http.HandleFunc("/url3", responseString)
	log.Println("Listening:", port[1:])
	log.Fatal(http.ListenAndServe(port, nil))
}

// handleURL - test with: curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url1
func responseText(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")
	decoder := json.NewDecoder(r.Body)

	var ev httpMessage
	errDecode := decoder.Decode(&ev)
	if errDecode != nil {
		log.Fatal("error decode: ", errDecode)
		return
	}
	log.Println("response: ", ev.Msg)
	w.Write([]byte(ev.Msg)) // echo as text
}

// handleURL - test with: curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url2
func responseJSON(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")
	decoder := json.NewDecoder(r.Body)

	var ev httpMessage
	errDecode := decoder.Decode(&ev)
	if errDecode != nil {
		log.Fatal("error decode: ", errDecode)
		return
	}
	log.Println("response: ", ev.Msg)
	response := httpMessage{Msg: ev.Msg}
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response) // echo as JSON
}

// handleURL - test with: curl -X POST -H 'Content-Type: application/json' -d "{\"message\": \"Hi\"}" http://127.0.0.1:8080/url3
func responseString(w http.ResponseWriter, r *http.Request) {
	log.Println("new request ...")
	decoder := json.NewDecoder(r.Body)

	var ev httpMessage
	errDecode := decoder.Decode(&ev)
	if errDecode != nil {
		log.Fatal("error decode: ", errDecode)
		return
	}
	log.Println("response: ", ev.Msg)

	var b bytes.Buffer
	response := httpMessage{Msg: ev.Msg}
	json2stream := json.NewEncoder(&b)
	json2stream.Encode(&response)

	s := b.String()
	w.Write([]byte(s)) // echo as text in JSON format
}
