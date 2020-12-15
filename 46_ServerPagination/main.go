package main

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

const theRoute = "/"
const theURL = "127.0.0.1:8080"
const version = "0.0.1"

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/css/", HomeHandler)
	r.HandleFunc("/posts/{pageNumber}/{itemsPage:[0-9]+}", PostsHandler)

	log.Println("alive on", theURL)
	http.Handle("/", r)
	http.ListenAndServe(theURL, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("home", version, r.URL)
	http.ServeFile(w, r, r.URL.Path[1:])
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {

	requestParams := mux.Vars(r)
	pageNumber := requestParams["pageNumber"]
	itemsPerPage := requestParams["itemsPage"]

	log.Println("posts call: ", pageNumber, itemsPerPage)

	theJSON, err := ComposeResponse(pageNumber, itemsPerPage, GetDBCredentials())
	if err != nil {
		log.Fatal("PostsHandler.ComposeResponse:", err)
	}

	log.Println(theJSON)

	json2stream := json.NewEncoder(w)
	json2stream.Encode(&theJSON)
}

func GetDBCredentials() *Credentials {

	DBCredentials := new(Credentials)
	DBCredentials.Host = "192.168.1.21"
	DBCredentials.User = "postgres"
	DBCredentials.Pass = "admin"
	DBCredentials.Port = "5432"
	DBCredentials.DatabaseName = "tests"
	DBCredentials.DatabaseTable = ""
	DBCredentials.RDBMS = "postgres"

	return DBCredentials
}
