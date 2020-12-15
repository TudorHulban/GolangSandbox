package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

// restart broker before testing
func TestRegisterConsumer(t *testing.T) {
	client := &http.Client{}
	jsonString := []byte(`{"code": "C1", "socket": "127.0.0.55:8085"}`)
	route := urlRegisterConsumer

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = route
	apiURLFormatted := u.String()

	request, err := http.NewRequest("POST", apiURLFormatted, bytes.NewBuffer(jsonString))
	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}
	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}
	r, _ := readResponse(response)
	if r != "ok" {
		t.Error("producer not registered correctly", r)
	}
}

func TestConsumeEmptyQueue(t *testing.T) {
	client := &http.Client{}
	jsonString := []byte(`{"code": "C1"}`)
	route := urlConsumeEvent

	u, _ := url.ParseRequestURI(apiURL)
	u.Path = route
	apiURLFormatted := u.String()

	request, err := http.NewRequest("POST", apiURLFormatted, bytes.NewBuffer(jsonString))
	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Error("NewRequest constructor error: ", err)
	}
	response, err := client.Do(request)
	if err != nil {
		t.Error("NewRequest  error: ", err)
	}
	r, _ := readResponse(response)
	if r != "no work" {
		t.Error("error consume: ", r)
	}
}
