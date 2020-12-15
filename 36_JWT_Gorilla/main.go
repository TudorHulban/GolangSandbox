package main

//check also https://dzone.com/articles/authentication-in-golang-with-jwts

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const HTTPPort = 3000
const URLLogin = "/login"
const URLRestricted = "/a"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HandlerPublic)
	r.HandleFunc("/login", HandlerLogin)
	r.Handle("/a", authMiddleware(hAdmin))

	r.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r)))
}

func HandlerPublic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Public Route")
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Route")
}

var hAdmin = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Admin Route")
})

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
