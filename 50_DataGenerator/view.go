package datageneration

import (
	"os"
	"strconv"
	"strings"
)

// DataToFile - exports generated data to file
func DataToFile(writeTo string, data *GenData) error {
	file, err := os.Create(writeTo)
	if err != nil {
		return err
	}
	defer file.Close()

	// write header
	fileHeader := []string{}
	for _, columnName := range data.ColumnNames {
		fileHeader = append(fileHeader, strings.Trim(columnName, " "))
	}

	file.WriteString(strings.Join(fileHeader, ",") + "\n")
	file.Sync()

	// write content
	for _, rowVals := range data.Rows {
		fileRow := []string{}

		for _, value := range rowVals {

			switch value.(type) {
			case string:
				fileRow = append(fileRow, value.(string))
			case int:
				fileRow = append(fileRow, strconv.FormatInt(int64(value.(int)), 10))
			}
		}
		file.WriteString(strings.Join(fileRow, ",") + "\n")
	}
	file.Sync()
	return nil
}
