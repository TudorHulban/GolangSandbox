package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"encoding/json"

	"github.com/gorilla/mux"
)

const theURL = "127.0.0.1:8000"
const theItem = "tasks"

/* {"items":[{"name":"item a", "description":"a"},{"name":"item b", "description":"b"}]}  */

type item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type items struct {
	Items []item `json:items`
}

var theItems items

func main() {
	theItems.Items = append(theItems.Items, item{"new", "x"})
	theItems.Items = append(theItems.Items, item{"new", "y"})
	fmt.Println("items:", theItems)

	r := mux.NewRouter()
	r.HandleFunc("/"+theItem, collectionFetch).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         theURL,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Marshal repo and send it to http serve
func collectionFetch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("collectionFetch items:", theItems)
	results, err := json.Marshal(theItems)
	if err != nil {
		log.Println("json.Marshal: ", err)
	}

	fmt.Println("collectionFetch results: ", string(results[:]))
	serveJSON(w, string(results))
}

func serveError(w http.ResponseWriter, theError string) {
	serveJSON(w, "error -"+theError)
}

func serveJSON(w http.ResponseWriter, content string) {
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&content)
}


//https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
//https://medium.com/@kelvin_sp/building-and-testing-a-rest-api-in-golang-using-gorilla-mux-and-mysql-1f0518818ff6
