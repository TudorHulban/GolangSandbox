package datageneration
import (
	"os"
	"strconv"
	"strings"
)

// DataToFile - exports generated data to file
func DataToFile(pWriteTo string, pData *GenData) error {
	hFile, err := os.Create(pWriteTo)
	if err != nil {
		return err
	}
	defer hFile.Close()

	// write header
	fileHeader := []string{}
	for _, columnName := range pData.ColumnNames {
		fileHeader = append(fileHeader, strings.Trim(columnName, " "))
	}

	hFile.WriteString(strings.Join(fileHeader, ",") + "\n")
	hFile.Sync()

	// write content
	for _, rowVals := range pData.Rows {
		fileRow := []string{}

		for _, value := range rowVals {

			switch value.(type) {
			case string:
				fileRow = append(fileRow, value.(string))
			case int:
				fileRow = append(fileRow, strconv.FormatInt(int64(value.(int)), 10))
			}
		}
		hFile.WriteString(strings.Join(fileRow, ",") + "\n")
	}
	hFile.Sync()
	return nil
}