package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type FieldConfig struct {
	Name         string `json:"name"`
	Type         int    `json:"ftype"`
	Length       int    `json:"flength"`
	PositiveOnly bool   `json:"fpositive"`
	MinValue     int    `json:"min"`
	MaxValue     int    `json:"max"`
}

type GenConfig struct {
	Configuration []FieldConfig `json:"config"`
	NoRows        int64         `json:"norows"`
}

type GenData struct {
	ColumnNames []string        `json:"columns"`
	Rows        [][]interface{} `json:"rowsdata"`
}

func main() {
	x := NewConfig(4)
	err := DoPost("http://localhost:3002", x)
	log.Println("DoPost: ", err)
}

func DoPost(pURL string, pConfig *GenConfig) error {
	u, _ := url.ParseRequestURI(pURL)
	apiURLFormatted := u.String()

	client := &http.Client{}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(pConfig)
	req, err := http.NewRequest("POST", apiURLFormatted, buf)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	//log.Println("body:", string(body))

	var data GenData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	log.Println(data.ColumnNames)
	log.Println(data.Rows)
	return err
}

func NewConfig(pNoRows int64) *GenConfig {
	instance := new(GenConfig)
	instance.NoRows = pNoRows
	instance.Configuration = make([]FieldConfig, 2)

	instance.Configuration[0].Name = "a"
	instance.Configuration[0].Type = 2
	instance.Configuration[0].Length = 5
	instance.Configuration[0].MinValue = 0
	instance.Configuration[0].MaxValue = 100
	instance.Configuration[0].PositiveOnly = true

	instance.Configuration[1].Name = "b"
	instance.Configuration[1].Type = 0
	instance.Configuration[1].Length = 5
	instance.Configuration[1].MinValue = 0
	instance.Configuration[1].MaxValue = 100
	instance.Configuration[1].PositiveOnly = true
	return instance
}
