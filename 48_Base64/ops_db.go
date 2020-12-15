package main

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type File struct {
	ID      int64
	Name    string
	Content string
}

var connDb *pg.DB

func init() {
	connDb = pg.Connect(&pg.Options{
		//Addr:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "pp",
		Database: "postgres",
	})
	log.Println(createTable4Model(interface{}(&File{})))
	errDelete := os.Remove(targetPath)
	log.Println("errDelete:", errDelete)
	log.Println("init was run")
}

func createTable4Model(pModel interface{}) error {
	return connDb.CreateTable(pModel, &orm.CreateTableOptions{Temp: false, IfNotExists: true, FKConstraints: true})
}

func insertFile(pFile *File) error {
	return connDb.Insert(pFile)
}

func fetchFile(pID int64) (*File, error) {
	result := new(File)
	result.ID = pID
	errFetch := connDb.Select(result)
	return result, errFetch
}
