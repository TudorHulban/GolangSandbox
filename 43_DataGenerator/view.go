package main

import (
	"os"
	"strconv"
	"strings"
)

func DataToFile(pWriteTo string, pData *GenData) error {
	hFile, err := os.Create(pWriteTo)
	if err != nil {
		return err
	}
	defer hFile.Close()

	// write header
	fileHeader := []string{}
	for _, v := range pData.ColumnNames {
		fileHeader = append(fileHeader, strings.Trim(v, " "))
	}

	hFile.WriteString(strings.Join(fileHeader, ",") + "\n")
	hFile.Sync()

	// write content
	for _, rowVals := range pData.Rows {

		fileRow := []string{}
		for _, v := range rowVals {

			switch v.(type) {
			case string:
				fileRow = append(fileRow, v.(string))
			case int:
				fileRow = append(fileRow, strconv.FormatInt(int64(v.(int)), 10))
			}
		}
		hFile.WriteString(strings.Join(fileRow, ",") + "\n")
	}
	hFile.Sync()
	return nil
}
