package main

import (
	"log"
	"testing"
)

func TestRowsInsert(t *testing.T) {
	h := new(DBServerInfo)
	h.Host = "192.168.1.13"
	h.User = "root"
	h.Pass = "root"
	h.Port = "3306"
	h.DatabaseName = "teste"
	h.DatabaseTable = "target"
	h.RDBMS = "mysql"

	r := new(JSON4Table)
	slice2Inject := []byte(`{"Data":[{"id":"1","name":"john"},{"id":"2","name":"popescu"}],"Database":"teste","Table":"target"}`)

	rowsInserted, err := r.Insert(h, slice2Inject)
	if err != nil {
		t.Error(err)
	}
	log.Println("rowsInserted:", rowsInserted)
}
