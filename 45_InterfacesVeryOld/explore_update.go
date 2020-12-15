package main

import (
	"log"
)

//SQLUpdate - used for testing only the sql update
func SQLUpdate(serverInfo *DBServerInfo) {
	columnTarget := "name"
	columnClause := "id"

	var inject2PreparedStatement []updateFieldValues

	var upd1 updateFieldValues
	upd1.clauseColumn = "1"
	upd1.replaceWith = "x1"

	var upd2 updateFieldValues
	upd2.clauseColumn = "2"
	upd2.replaceWith = "y1"

	inject2PreparedStatement = append(inject2PreparedStatement, upd1)
	inject2PreparedStatement = append(inject2PreparedStatement, upd2)

	log.Println("inject2PreparedStatement:", inject2PreparedStatement)

	statement := "update teste.accounts set " + columnTarget + "=?" + " where " + columnClause + "=?"
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

	for key, valuesRow := range inject2PreparedStatement {
		log.Println("replaceWith:", valuesRow.replaceWith, " - where:", valuesRow.clauseColumn)
		log.Println("update teste.accounts set " + columnTarget + "=" + valuesRow.replaceWith + " where " + columnClause + "=" + valuesRow.clauseColumn)

		sqlResult, err := dml.Exec(valuesRow.replaceWith, valuesRow.clauseColumn)
		if err != nil {
			log.Fatal("dml.Exec:", err)
		}

		updated, err := sqlResult.RowsAffected()
		if err != nil {
			dbTransaction.Rollback()
			log.Println("sqlResult.RowsAffected:", err)
			break
		}
		log.Println(key, "updated:", updated)
	}
	dbTransaction.Commit()
}
