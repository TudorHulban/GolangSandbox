package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"time"
)

const port = 8080

type httpResponse struct {
	Msg string `json:"message"`
}

func main() {
	http.HandleFunc("/x", handlerOne)
	http.HandleFunc("/y", handlerTwo)

	http.ListenAndServe((fmt.Sprintf(":%v", port)), nil)
}

func handlerOne(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	response := httpResponse{Msg: "xxxxx"}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	fmt.Fprint(w, string(data))
	fmt.Println("duration:", time.Now().Sub(startTime))
}

func handlerTwo(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	response := httpResponse{Msg: "yyyyy"}
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)

	fmt.Println("duration:", time.Now().Sub(startTime))
}
