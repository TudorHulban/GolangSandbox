package main

import (
	"database/sql"
	"log"
)

//DBServer - struct implementing methods for interfaces related to database operations
type DBServer struct {
	dbHandler *sql.DB
}

//DBRead - interface. satisfied by functions that receive SQL query and return pointers to rows and error
type DBRead interface {
	Read(command string) (*sql.Rows, error)
}

//Read - satisfies DBRead interface. receives a SQL query and returns pointers to rows and error
func (db *DBServer) Read(command string) (*sql.Rows, error) {
	rows, err := db.dbHandler.Query(command)
	if err != nil {
		log.Println("this.dbHandler.Query:", err)
		return nil, err
	}
	return rows, nil
}

//DBPing - interface. satisifed by functions that set a database handler and return the error when performing a database Ping
type DBPing interface {
	Ping() error
}

//Ping - satisfies DBPing interface. based on set handler returns error when performing a database ping
func (db *DBServer) Ping() error {
	return db.dbHandler.Ping()
}
