package main

import (
	"log"
	"testing"
)

func TestJSONUpdate(t *testing.T) {

	h := new(DBServerInfo)
	h.Host = "192.168.1.13"
	h.User = "root"
	h.Pass = "root"
	h.Port = "3306"
	h.DatabaseName = "teste"
	h.RDBMS = "mysql"

	r := new(JSON4Update)
	//update - column to update , where - column for clause
	slice2Inject := []byte(`{"Data":[{"Update":{"name":"Tudor"},"Where":{"id":"1"}},{"Update":{"name":"Franc"},"Where":{"id":"2"}}],"Database":"teste","Table":"accounts"}`)

	rowsAffected, err := r.Update(h, slice2Inject)
	if err != nil {
		t.Error(err)
	}

	log.Println("rowsAffected:", rowsAffected)
}
