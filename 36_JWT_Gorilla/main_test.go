package main

import (
	"encoding/json"
	"fmt"
	_ "io"
	_ "io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

var token string

type JwtToken struct {
	Token string `json:"token"`
}

func TestHandleLoginReq(t *testing.T) {
	r, err := http.PostForm("http://localhost:"+strconv.Itoa(HTTPPort)+URLLogin, url.Values{"username": {"x"}, "password": {"y"}})
	if err != nil {
		t.Error("Post Form error: ", err)
	}
	defer r.Body.Close()

	var tok JwtToken
	if err = json.NewDecoder(r.Body).Decode(&tok); err != nil {
		fmt.Println("new decoder: ", err)
	}

	token = tok.Token
	fmt.Println("token: ", tok.Token)
}

func TestHandleRestrictedRoutes(t *testing.T) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://localhost:"+strconv.Itoa(HTTPPort)+URLRestricted, nil)
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}

	if len(token) == 0 {
		token = "xxx"
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Accept", "application/json")

	fmt.Println("request: ", request)

	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}

	fmt.Println("response: ", response)
}
