package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

func getAsJSON(db *sql.DB, sql string) (string, error) {
	stmt, errPrepare := db.Prepare(sql)
	if errPrepare != nil {
		return "", errors.WithMessagef(errPrepare, "could not prepare: %v", sql)
	}
	defer stmt.Close()

	rows, errQuery := stmt.Query()
	if errQuery != nil {
		return "", errors.WithMessage(errQuery, "could not query")
	}
	defer rows.Close()

	columns, errCols := rows.Columns()
	if errCols != nil {
		return "", errors.WithMessage(errCols, "could not extract columns")
	}
	queryData := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(columns))

	//prepares memory addresses of the pointers to the data
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		errScan := rows.Scan(scanArgs...) //fills the addresses with data
		if errScan != nil {
			return "", errors.WithMessage(errScan, "could not scan row")
		}

		rowData := make(map[string]interface{})

		for ix, col := range columns {
			columnValue := values[ix]

			byteValue, isByte := columnValue.([]byte)
			if isByte {
				rowData[col] = string(byteValue)
				continue
			}
			rowData[col] = columnValue
		}
		queryData = append(queryData, rowData)
	}

	//wrap up results
	finalData := map[string][]map[string]interface{}{"rows": queryData}
	log.Println("query data:", finalData)
	jsonData, err := json.Marshal(finalData)
	if err != nil {
		return "", err
	}
	log.Println("marshaled data:", string(jsonData[:]))

	return string(jsonData), nil
}
