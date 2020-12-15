package main

import (
	"encoding/json"
	"log"
	"strings"
)

//DBServerInfo - struct. fields necessary to connect to database server
type DBServerInfo struct {
	Host          string
	Port          string
	User          string
	Pass          string
	RDBMS         string //type of database. values are:
	DatabaseName  string
	DatabaseTable string
}

//JSON4Table - struct implementing receiver for interface JSON2Table
type JSON4Table struct{}

//JSON2Table - interface. satisfied by functions receiving database server info and raw JSON information. JSON info contains database table where to insert
type JSON2Table interface {
	Insert(serverInfo *DBServerInfo, rawJSONData []byte) (int64, error)
}

//Insert - method satisfying interface JSON2Table. takes database server info and JSON information and pushes it to table specified in JSON
func (j4t *JSON4Table) Insert(serverInfo *DBServerInfo, rawJSONData []byte) (int64, error) {
	columnNames := make([]string, 0)
	var results map[string]interface{}

	if err := json.Unmarshal(rawJSONData, &results); err != nil {
		log.Fatal("json.Unmarshal:", err)
	}

	log.Println("results:", results)

	databaseName := results["Database"].(string)
	tableName := results["Table"].(string)
	rowData := results["Data"].([]interface{})

	jsonData := rowData[0]

	row := jsonData.(map[string]interface{})
	preparedValues := make([]string, 0)

	for k := range row {
		columnNames = append(columnNames, k)
		preparedValues = append(preparedValues, "?")
	}

	log.Println("columnNames:", columnNames)
	statement := "insert into " + databaseName + "." + tableName + "(" + strings.Join(columnNames, ",") + ")" + " values(" + strings.Join(preparedValues, ",") + ")"

	log.Println("statement:", statement)

	db, err := new(DBHandler).Connect(serverInfo)
	if err != nil {
		log.Fatal("new(DBHandler).Connect:", err)
	}

	dbTransaction, err := db.Begin()
	if err != nil {
		log.Fatal("db.Begin:", err)
	}

	dml, err := dbTransaction.Prepare(statement)
	if err != nil {
		log.Fatal("db.Prepare:", err)
	}
	defer dml.Close()

	log.Println("dml:", dml)

	var rowCount int64

	//needs sort, gives error sometimes on columns order

	for _, rowInfo := range rowData {
		log.Println("rowInfo:", rowInfo)

		values := make([]interface{}, 0)

		for _, columnValue := range rowInfo.(map[string]interface{}) {
			values = append(values, columnValue)
		}

		log.Println("values:", values)

		sqlResult, err := dml.Exec(values...)
		if err != nil {
			log.Fatal("dml.Exec:", err)
		}

		inserted, err := sqlResult.RowsAffected()
		if err != nil {
			log.Fatal("sqlResult.RowsAffected:", err)
		}

		rowCount = rowCount + inserted
	}

	dbTransaction.Commit()
	return rowCount, nil
}
