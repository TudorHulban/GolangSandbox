package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type jsonAnswer struct {
	Response string `json:"response"`
}

type request struct {
	code    string // producer code
	id      int64
	ttl     int64
	payload []string
}

func registerWithBroker() (string, error) {
	json := `{"code": "` + producerCode + `", "ip":"` + ip + `"}`
	//log.Println("json:", json)
	return jsonRequest(brokerAPIurl, routeRegister, json, "POST")
}

// postEvent - to refactor - later
func postRequest(pRequest request) (string, error) {
	json := `{"code": "` + pRequest.code + `", "id":` + strconv.FormatInt(pRequest.id, 10) + `,"ttl":` + strconv.FormatInt(pRequest.ttl, 10) + `,"payload":["` + strings.Join(pRequest.payload[:], "\", \"") + `"]}`
	log.Println("json:", json)
	return jsonRequest(brokerAPIurl, routeEvents, json, "POST")
}

func readyRequest(pRequest request) bool {
	json := `{"code": "` + pRequest.code + `", "id":` + strconv.FormatInt(pRequest.id, 10) + `}`
	log.Println("json:", json)
	ready, errStatus := jsonRequest(brokerAPIurl, routeStatus, json, "POST")
	log.Println(ready, errStatus)
	return false
}

func jsonRequest(pAPIurl, pRoute, pJSON, pHTTPMethod string) (string, error) {
	client := &http.Client{}
	jsonString := []byte(pJSON)

	u, _ := url.ParseRequestURI(pAPIurl)
	u.Path = pRoute
	apiURLFormatted := u.String()

	request, err := http.NewRequest(pHTTPMethod, apiURLFormatted, bytes.NewBuffer(jsonString))
	request.Header.Set("X-Custom-Header", "myvalue")
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("1")
		return "", err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println("2")
		return "", err
	}
	return readResponse(response)
}

func readResponse(pResponse *http.Response) (string, error) {
	defer pResponse.Body.Close()
	body, errRead := ioutil.ReadAll(pResponse.Body)
	if errRead != nil {
		return "", errRead
	}

	var result jsonAnswer
	errConvert := json.Unmarshal(body, &result)
	return result.Response, errConvert
}
