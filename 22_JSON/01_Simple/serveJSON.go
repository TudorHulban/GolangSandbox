package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Signal Type for representing one signal.
type Signal struct {
	Value int `json:"value"`
}

// Signals Type for sending signals.
type Signals []Signal

const theRoute = "/"
const theURL = "localhost:8080"

func main() {
	http.HandleFunc(theRoute, handleRoutes)
	panic(http.ListenAndServe(theURL, nil))
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" {
		loginHandler(w, r)
		return
	}
	if r.URL.Path == "/jarray" {
		jarrayHandler(w, r)
		return
	}
	if r.URL.Path == "/darray" {
		arrayHandler(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("login call")
	theJSON := []byte(`{"Response":{"Status":"success","StatusCode":200,"Failed":false,"Message":"success"}}`)

	w.Header().Set("Content-Type", "application/json")
	w.Write(theJSON)
}

func jarrayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("json array")
	theJSON := []byte(`{"response":[{"t1":1},{"t1":2}]}`)

	w.Header().Set("Content-Type", "application/json")
	w.Write(theJSON)
}

func arrayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("defined json array")
	signals := Signals{
		Signal{
			Value: 101,
		},
		Signal{
			Value: 78,
		},
	}
	theJSON, errMa := json.Marshal(signals)
	if errMa != nil {
		w.Write([]byte(errMa.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(theJSON)
}
