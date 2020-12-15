package main

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteClient struct {
	pathDB string
	connDB *sql.DB
}

func NewSQLite(path string) (SQLiteClient, error) {
	if path == "" {
		return SQLiteClient{}, errors.New("invalid database path")
	}
	db, errCo := sql.Open("sqlite3", path)
	return SQLiteClient{
		pathDB: path,
		connDB: db,
	}, errCo
}

func (s SQLiteClient) CreateTable(tableName, ddl string) error {
	_, errExec := s.connDB.Exec("CREATE TABLE " + tableName + ddl)
	return errExec
}

// TableExists returns true if table exists.
func (s SQLiteClient) TableExists(tableName string) bool {
	var occurences int
	_ = db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", tableName).Scan(&occurences)

	return (occurences == 1)
}
