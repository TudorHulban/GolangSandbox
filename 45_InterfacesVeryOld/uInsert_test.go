package main

import (
	"testing"
)

func TestJSONInsert(t *testing.T) {
	insert := `{"Database":"teste","Table":"accounts","Data":[{"id":"1","name":"x"}]}`

	jRaw := new(JSONRAW)
	jRaw.rawJSONData = []byte(insert)

	if err := jRaw.Insert("mysql", "teste", "accounts"); err != nil {
		t.Error(err)
	}
}
