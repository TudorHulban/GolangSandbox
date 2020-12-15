package main

import (
	"database/sql"
	"log"
)

//SQLToWriter - structure composition. aggregates 2 interfaces
type SQLToWriter struct {
	Reader DBRead
	Writer DBWrite
}

//Execute - method. orphan?
func (receiver *SQLToWriter) Execute(sqlCommand string, targetPath string) error {
	rows, err := receiver.Reader.Read(sqlCommand) //see http://go-database-sql.org/retrieving.html
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return err
	}

	columnNames, err := rows.Columns()
	if err != nil {
		log.Fatal("rows.Columns:", err)
	}

	columnCount := len(columnNames)
	rowData := make([]interface{}, columnCount) //see http://go-database-sql.org/varcols.html
	rowValues := make([]sql.RawBytes, columnCount)

	for i := range rowValues {
		rowData[i] = &rowValues[i]
	}

	var rowID int64
	var columnValueForCurrentRow string
	queryResults := make(map[int64][]string)

	for i := range columnNames {
		queryResults[0] = append(queryResults[0], columnNames[i])
	}

	for rows.Next() {
		err := rows.Scan(rowData...)
		if err != nil {
			log.Fatal("rows.Scan:", err)
		}

		for _, col := range rowValues {

			if col == nil {
				columnValueForCurrentRow = "NULL"
			} else {
				columnValueForCurrentRow = string(col)
			}
			log.Println(columnValueForCurrentRow)
			queryResults[rowID+1] = append(queryResults[rowID+1], columnValueForCurrentRow)
		}

		rowID++
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("rows.Next:", err)
	}

	receiver.Writer.SetContent(queryResults)
	return receiver.Writer.Write(targetPath)
}

func newSQLReadWriter(r DBRead, w DBWrite) *SQLToWriter {
	rw := new(SQLToWriter)
	rw.Reader = r
	rw.Writer = w

	return rw
}
