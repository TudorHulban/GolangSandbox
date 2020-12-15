package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type httpResponse struct {
	Msg string `json:"message"`
}

const ( // iota is reset to 0
	numeric    = iota // == 0
	characters = iota // == 1
	id         = iota // == 2
)

const port = 3002

const values = `{"fields":[{"name":"a"}, {"name":"b"}, {"name":"c"}], "values":[{"f1":"1","f2":"2","f3":"c"}]}`
const config = `{"config":[{"name":"a","ftype":0, "flength":5, "fpositive":true, "min":5, "max":10000},{"name":"b","ftype":1, "flength":4, "fpositive":true}], "norows":5}`

func main() {
	http.HandleFunc("/", dataHandler)
	http.ListenAndServe((fmt.Sprintf(":%v", port)), nil)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	bodyString := ioutil.NopCloser(bytes.NewBuffer(body))
	log.Println("body:", bodyString)

	buf := new(bytes.Buffer)
	buf.ReadFrom(bodyString)

	configResponse, err := GetConfig(buf.String())
	if err != nil {
		return
	}
	configData, err := GetData(configResponse)
	if err != nil {
		return
	}

	//log.Println("column names: ", configData.ColumnNames)
	log.Println("rows: ", configData.Rows)

	response := GenData{configData.ColumnNames, configData.Rows}
	w.Header().Set("Content-Type", "application/json")
	json2stream := json.NewEncoder(w)
	json2stream.Encode(&response)

	duration := time.Now().Sub(startTime)
	fmt.Println("y:", duration)
}
