package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// no router.

// Signal Type for representing one signal.
type Signal struct {
	Value int `json:"value"`
}

// Signals Type for sending signals.
type Signals []Signal

const (
	theRoute = "/"
	theURL   = "localhost:8080"
)

func main() {
	log.Println("starting...")

	http.HandleFunc(theRoute, handleRoutes)
	panic(http.ListenAndServe(theURL, nil))
}

func handleRoutes(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/login" {
		loginHandler(w, r)
		return
	}

	if r.URL.Path == "/array" {
		arrayHandler(w, r)
		return
	}

	http.ServeFile(w, r, r.URL.Path[1:])
}

// loginHandler Handler with no marshalling, just returning slice of bytes.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`[{"value":101},{"value":78}]`))
}

func arrayHandler(w http.ResponseWriter, r *http.Request) {
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
