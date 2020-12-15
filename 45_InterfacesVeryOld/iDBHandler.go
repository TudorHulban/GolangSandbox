package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DBHandler - struct implementing methods for interface
type DBHandler struct{}

//ConnectDB - interface. satisfied by functions that receive database server information and return a database handler and error
type ConnectDB interface {
	Connect(server *DBServerInfo) (*sql.DB, error) //with sql.Open performed. ready to use
}

//Connect -  method satisfying interface, takes database server information and returns a database handler and error
func (h *DBHandler) Connect(server *DBServerInfo) (*sql.DB, error) {
	loginString := server.User + ":" + server.Pass + "@tcp(" + server.Host + ":" + server.Port + ")/" + server.DatabaseName
	log.Println("loginString:", loginString)

	return sql.Open(server.RDBMS, loginString)
}
