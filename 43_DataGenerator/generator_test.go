package main

import (
	"log"
	"testing"
)

//const testJSON = `{"config":[{"name":"a","ftype":0, "flength":5, "fpositive":true, "min":5, "max":10000},{"name":"b","ftype":1, "flength":4, "fpositive":true}], "norows":5}`
const testJSON = `{"config":[{"name":"id","ftype":2},{"name":"a","ftype":0, "flength":5, "fpositive":true, "min":1, "max":100}], "norows":4}`

func TestController(t *testing.T) {

	testConfig, err := GetConfig(testJSON)
	if err != nil {
		t.Error("GetConfig: ", err)
	}
	log.Println("GetConfig: ", *testConfig)

	generData, err := GetData(testConfig)
	if err != nil {
		t.Error("GetData: ", err)
	}
	log.Println("GetData: ", *generData)

	err = DataToFile("xxx.csv", generData)
	if err != nil {
		t.Error("DataToFile: ", err)
	}
}
