package main

import (
	"io/ioutil"
	"log"
)

func main() {
	h := new(DBServerInfo)
	h.Host = "192.168.1.13"
	h.User = "root"
	h.Pass = "root"
	h.Port = "3306"
	h.DatabaseName = "teste"
	h.RDBMS = "mysql"

	connect := DBHandler{}
	db, err := connect.Connect(h)
	if err != nil {
		log.Fatal("DBHandler.sql.Open:", err)
	}

	p := new(DBServer)
	p.dbHandler = db

	err = p.Ping()
	if err != nil {
		log.Fatal("p.Ping:", err)
	}

	//r := new(RowsInfoToFile)

	r := new(RowsInfoToJSON)

	rw := newSQLReadWriter(p, r)
	rw.Execute("select * from accounts", "z:\\sqlresults.txt")

	j2 := new(JSON4DB)

	rawJSONData, err := ioutil.ReadFile("z:\\sqlresults.txt")
	if err != nil {
		log.Fatal("ioutil.ReadFile:", err)
	}

	j2.Insert("MySQL", "teste", "accounts", rawJSONData)
}
