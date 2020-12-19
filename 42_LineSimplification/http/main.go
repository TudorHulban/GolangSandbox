package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

const noRows = 20

func main() {
	x := NewConfig(noRows)
	err := DoPost("http://localhost:3002", x)
	log.Println("DoPost error: ", err)
}

func DoPost(url string, cfg *GenConfig) error {
	u, _ := url.ParseRequestURI(url)
	apiURLFormatted := u.String()

	client := &http.Client{}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(cfg)

	req, errReq := http.NewRequest("POST", apiURLFormatted, buf)
	if errReq != nil {
		return errReq
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var data GenData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	log.Println(data.ColumnNames)
	log.Println("points before: ", data.Rows)

	rows := data.Rows
	floats := SliceInterfaceToFloats(&rows)
	optimPoints := Simplify(*floats, 2, true)
	log.Println("points after: ", optimPoints)
	return err
}

func NewConfig(noRows int64) *GenConfig {
	f1 := FieldConfig{
		Name:         "a",
		Type:         2,
		Length:       5,
		MinValue:     0,
		MaxValue:     100,
		PositiveOnly: true,
	}

	f2 := FieldConfig{
		Name:         "b",
		Type:         0,
		Length:       5,
		MinValue:     0,
		MaxValue:     100,
		PositiveOnly: true,
	}

	return &GenConfig{
		NoRows:        noRows,
		Configuration: []FieldConfig{f1, f2},
	}
}

func SliceInterfaceToFloats(s *[][]interface{}) *[][]float64 {
	instance := make([][]float64, 0)

	for _, v := range *s {
		row := make([]float64, 0)

		for _, vv := range v {
			row = append(row, vv.(float64))
		}
		instance = append(instance, row)
	}

	return &instance
}
