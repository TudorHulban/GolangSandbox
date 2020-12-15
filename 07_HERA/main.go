package main

import (
	"fmt"

	"time"
)

func main() {
	s := DBSQLiteInfo{
		pathSQLiteFile: "./test.dbf",
	}

	initDB(s, "")

	/*
		var p DBPostgresInfo
		p.ip = "192.168.1.3"
		p.port = 5432
		p.user = "golang"
		p.password = "golang"

		initDB(p, "db_golang")
	*/

	/*
		var m DBMariaInfo
		m.ip = "192.168.1.12" //in container c0702c
		m.port = 3306
		m.user = "hera"
		m.password = "hera"

		initDB(m, "devops")
	*/
}

type rowValues struct {
	values []string
}

type tableDDL struct {
	tableName             string
	tableFields           string
	ColumnPKAutoincrement int
}

func initDB(rdbms RDBMS, dbName string) {
	db, err := rdbms.Open(dbName)
	checkErr(err, "initDB")

	var T1 = tableDDL{"version", prepareDBField("'name' text, 'date' text, 'description' text"), -1}
	var T2 = tableDDL{"dbtypes", prepareDBField("'code' text, 'description' text"), 1}
	var T3 = tableDDL{"checks", prepareDBField("'name' text, 'description' text, 'action' text, 'mins2run' integer, 'hourfrom' integer, 'hourto' integer, 'dayofweek' integer, 'dayofmonth' integer, 'lastrun' integer, 'enabled' text, 'systemid' integer"), 1}

	var schemaDefinition = []tableDDL{T1, T2, T3}

	for i := range schemaDefinition {
		tbCreated := true

		if !rdbms.TableExists(db, dbName, schemaDefinition[i].tableName) {
			fmt.Println("creating table: ", schemaDefinition[i].tableName)
			tbCreated = rdbms.CreateTable(db, dbName, schemaDefinition[i].tableName, schemaDefinition[i].tableFields, schemaDefinition[i].ColumnPKAutoincrement)
			fmt.Println("created table "+schemaDefinition[i].tableName+":", tbCreated)
			return
		}

		fmt.Println("already exists table: " + schemaDefinition[i].tableName)
	}

	//populate tables - test functions and if not different across RDBMSs take them out from interface
	var V1 = []string{"0.01", string(time.Now().AppendFormat(nil, "Mon Jan 02 2006 15:04:05 GMT-0700")[:]), "July 2018"}
	checkErr(rdbms.SingleInsert(db, "version", V1), "V1")

	var V2 = [][]string{[]string{"SQLITE", "Sqlite Connection"}, []string{"POSTGRES", "PostgreSQL Connection"}, []string{"ORACLE", "Oracle Connection"}}
	checkErr(rdbms.BulkInsert(db, "dbtypes", []string{"code", "description"}, V2), "V2")

	var V3 = [][]string{[]string{"test1", "test check", "action 1", "20", "0", "24", "127", "2147483647", "0", "Y", "1"}}
	checkErr(rdbms.BulkInsert(db, "checks", []string{"name", "description", "action", "mins2run", "hourfrom", "hourto", "dayofweek", "dayofmonth", "lastrun", "enabled", "systemid"}, V3), "V3")
}
