package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var tokenString string
var apiURL = "http://localhost:" + port

type JwtToken struct {
	Token string `json:"token"`
}

func TestIndexHandler(t *testing.T) {
	client := &http.Client{}
	route := URLIndex

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = route
	apiURLFormatted := u.String()

	request, err := http.NewRequest("GET", apiURLFormatted, nil)
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}

	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}
	log.Println("response: ", response)
}

func TestHandleLoginReq(t *testing.T) {
	client := &http.Client{}
	form := url.Values{"user": []string{"x"}, "password": []string{"y"}}

	route := URLLogin
	//apiURL := "http://requestbin.fullcontact.com"
	//route := "/10neo6o1/"

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = route
	apiURLFormatted := u.String()

	request, err := http.NewRequest("POST", apiURLFormatted, strings.NewReader(form.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}

	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}
	log.Println("response: ", response)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var tok JwtToken

	if err = json.Unmarshal(body, &tok); err != nil {
		log.Println("unmarshal:", err)
	}

	tokenString = tok.Token
	if len(tokenString) == 0 {
		t.Error("No Token.")
	} else {
		log.Println("token: ", tokenString)
	}
}

func TestHandleRestrictedRoutes(t *testing.T) {
	if len(tokenString) == 0 {
		t.Error("No Token.")
	} else {
		client := &http.Client{}
		request, err := http.NewRequest("GET", "http://localhost:"+port+URLRestricted, nil)

		if err != nil {
			t.Error("NewRequest constructor error: ", err)
		}

		request.Header.Set("Authorization", "Bearer "+tokenString)
		request.Header.Set("Accept", "application/json")
		log.Println("request: ", request)
		response, err := client.Do(request)

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		type Users struct {
			Users []string `json:"users"`
		}

		var um Users
		err = json.Unmarshal(body, &um)
		if err != nil {
			t.Error("Unmarshal  error: ", err)
		} else {
			log.Println("response: ", um.Users)
		}
	}
}
