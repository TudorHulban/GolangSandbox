package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Credentials struct {
	Host          string
	Port          string
	User          string
	Pass          string
	RDBMS         string //type of database. values are:
	DatabaseName  string
	DatabaseTable string
}

func Connect2RDBMS(dbInfo *Credentials) (*sql.DB, error) {
	loginString := "postgres://" + dbInfo.User + ":" + dbInfo.Pass + "@" + dbInfo.Host + ":" + dbInfo.Port + "/" + dbInfo.DatabaseName + "?sslmode=disable"

	log.Println("loginString:", loginString)

	db, err := sql.Open(dbInfo.RDBMS, loginString)
	if err != nil {
		log.Fatal("DBHandler.sql.Open:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DBHandler.Ping:", err)
	}

	return db, err
}
