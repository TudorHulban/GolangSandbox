package main

import (
	"context"
	"encoding/json"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	cache "github.com/TudorHulban/cachejwt"
)

const (
	port          = "3000"
	URLLogin      = "/login"
	URLRestricted = "/users"
	URLIndex      = "/"
)

const (
	secondsCacheExpiration = 600 //default value
	secondsJanitorClean    = 60
	secondsTokenExpiration = 600 //introduced as we could have different expiration values, ex. shorter for admin
)

var (
	c     = cache.NEWCache(secondsCacheExpiration)
	users = []string{"Joe", "Sam", "Mary"}
)

type httpResponse struct {
	Msg   string   `json:"message"` //with capital so it is exported
	Users []string `json:"users"`
}

type credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

func invalidMethod(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Invalid Method")
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Bad Request")
}

func serverError(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Server Error")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "dist/index.html")
	http.ServeFile(w, r, "index.html")
}

func loghin(w http.ResponseWriter, r *http.Request) {
	// curl -d "user=x&password=y" -X POST http://localhost:3000/login
	// check with https://requestbin.fullcontact.com/

	log.Println("------------------------")

	if r.Method != "POST" {
		invalidMethod(w, r)
		return
	}

	if err := r.ParseForm(); err != nil {
		invalidMethod(w, r)
		return
	}

	u := r.FormValue("user")
	p := r.FormValue("password")
	log.Println("credentials: ", u, p)

	userData := isGoodCredentials(u, p)
	if userData.ID == -1 { // TODO: change -1 to default nil value for type
		badRequest(w, r)
		return
	}

	if userData.ID != -1 {
		tok := new(JWTToken)

		item := cache.Item{Value: userData.ID, Expiration: 0}
		tokString, err := tok.New(secondsTokenExpiration, 1, "xxx", c, &item)
		if err != nil {
			serverError(w, r)
		} else {
			json.NewEncoder(w).Encode(tokenResponse{Token: tokString})
		}
	}
}

var getUsers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// curl -X GET -H "Authorization: Bearer <string token>" http://localhost:3000/users

	if r.Method != "GET" {
		invalidMethod(w, r)
		return
	}

	response := httpResponse{Msg: "xxxx", Users: users}
	json.NewEncoder(w).Encode(response)
})

func main() {
	log.Println("Creating Cache and Janitor")
	j := cache.NEWJanitor(c, secondsJanitorClean)
	j.Clean()

	log.Println("Running...")

	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("dist/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("dist/js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("dist/css"))))

	http.HandleFunc(URLLogin, loghin)
	http.Handle(URLRestricted, authMiddleware(getUsers, c))
	http.HandleFunc(URLIndex, IndexHandler)

	server := http.Server{Addr: ":" + port}

	go func() {
		log.Println("server started... ")

		err := server.ListenAndServe()
		if err != nil {
			log.Println("error: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit //blocking

	j.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("error: ", err)
	}
	log.Println("server stopped.")
}
