package datageneration

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

//const testJSON = `{"config":[{"name":"a","ftype":0, "flength":5, "fpositive":true, "min":5, "max":10000},{"name":"b","ftype":1, "flength":4, "fpositive":true}], "norows":5}`
const testJSON = `{"config":[{"name":"id","ftype":2},{"name":"a","ftype":0, "flength":5, "fpositive":true, "min":1, "max":100}], "norows":4}`

func TestController(t *testing.T) {
	testConfig, err := GetConfig(testJSON)
	assert.Nil(t, err)
	log.Println("GetConfig: ", *testConfig)

	generData, err := GetData(testConfig)
	assert.Nil(t, err)
	log.Println("GetData: ", *generData)

	assert.Nil(t, DataToFile("xxx.csv", generData))
}
