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
func (s2w *SQLToWriter) Execute(sqlCommand string, targetPath string) error {
	rows, err := s2w.Reader.Read(sqlCommand) //see http://go-database-sql.org/retrieving.html
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
		if err := rows.Scan(rowData...); err != nil {
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

	if err = rows.Err(); err != nil {
		log.Fatal("rows.Next:", err)
	}

	s2w.Writer.SetContent(queryResults)
	return s2w.Writer.Write(targetPath)
}

func newSQLReadWriter(r DBRead, w DBWrite) *SQLToWriter {
	return &SQLToWriter{
		Reader: r,
		Writer: w,
	}
}
