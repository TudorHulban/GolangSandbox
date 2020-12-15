package main

import (
	"encoding/json"
	"log"
	"strings"
)

//JSON4DB - struct implementing receiver for interface
type JSON4DB struct {
}

/*
JSON2DB - interface. satisfied by functions that receive type of database (ex. MySQL, PostgreSQL), database name, database table and raw JSON data for insertion.
JSON format is like: insert := `{"Database":"teste","Table":"accounts","Data":[{"id":"1","name":"x"}]}`
*/
type JSON2DB interface {
	Insert(dbType string, dbName string, dbTable string, rawJSONData []byte) error
}

//Insert - method, signature satisfies JSON2DB interface.
func (receiver *JSON4DB) Insert(dbType string, dbName string, dbTable string, rawJSONData []byte) error {

	var results map[string]interface{}

	err := json.Unmarshal(rawJSONData, &results)
	if err != nil {
		log.Fatal("json.Unmarshal:", err)
	}

	databaseName := results["Database"].(string)
	tableName := results["Table"].(string)
	rowData := results["Data"].([]interface{})

	log.Println("database:", databaseName)
	log.Println("table:", tableName)
	log.Println("rows:", rowData)

	insert := "insert into " + databaseName + "." + tableName + " "

	keys := make([]string, 0)
	values := make([]string, 0)

	for _, rowInfo := range rowData {

		row := rowInfo.(map[string]interface{})

		for mapKey, value := range row {
			keys = append(keys, string(mapKey))
			values = append(values, "'"+value.(string)+"'")
		}
	}

	insert = insert + "(" + strings.Join(keys, ",") + ")" + " values(" + strings.Join(values, ",") + ")"

	log.Println("insert:", insert)

	return nil
}
