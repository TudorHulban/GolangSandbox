package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const port = "8080"

var (
	users = []string{"Joe", "Sam", "Mary"}
)

type httpResponse struct {
	Msg   string   `json:"message"` //with capital so it is exported
	Users []string `json:"users"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	response := httpResponse{Msg: "xxxx", Users: users}
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Running...")

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("dist/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("dist/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("dist/css"))))
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/", indexHandler)

	server := http.Server{Addr: ":" + port}

	go func() {
		fmt.Println("started: ", time.Now().UTC())

		err := server.ListenAndServe()
		if err != nil {
			log.Println("error: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("error: ", err)
	}
	fmt.Println("server stopped", time.Now().UTC())
}
