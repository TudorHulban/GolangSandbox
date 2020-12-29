package main

import (
	"log"
	"net/http"
	"strings"

	cache "github.com/TudorHulban/cachejwt"
)

type UserData struct {
	ID        int64
	FirstName string
	LastName  string
}

func authMiddleware(next http.Handler, cache *cache.Cache) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware")

		tokenString := r.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		log.Println("token: ", tokenString)

		_, isTokenFound := cache.Get(tokenString)
		if isTokenFound {
			next.ServeHTTP(w, r)
		}
	})
}

// ValidateCredentials validates credentials against store
func isGoodCredentials(user, password string) *UserData {
	if user == "x" && password == "y" {
		return &UserData{
			ID:        1,
			FirstName: "John",
			LastName:  "Smith",
		}
	}
	return nil
}
