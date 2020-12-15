package main

import (
	_ "log"
	"testing"
)

func TestSQLUpdate(t *testing.T) {
	h := new(DBServerInfo)
	h.Host = "192.168.1.13"
	h.User = "root"
	h.Pass = "root"
	h.Port = "3306"
	h.DatabaseName = "teste"
	h.RDBMS = "mysql"

	SQLUpdate(h)
}
