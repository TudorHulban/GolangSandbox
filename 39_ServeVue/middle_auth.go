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

func authMiddleware(pNext http.Handler, pCache *cache.Cache) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware")

		tokenString := r.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		log.Println("token: ", tokenString)

		_, isTokenFound := pCache.Get(tokenString)
		if isTokenFound {
			pNext.ServeHTTP(w, r)
		}
	})
}

// ValidateCredentials validates credentials against store
func isGoodCredentials(pUser, pPassword string) *UserData {
	var u UserData

	if pUser == "x" && pPassword == "y" {
		u.ID = 1
		u.FirstName = "John"
		u.LastName = "Smith"
	} else {
		u.ID = -1
	}

	return &u
}
