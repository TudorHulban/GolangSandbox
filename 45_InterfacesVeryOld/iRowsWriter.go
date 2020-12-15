package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

//RowsInfoToFile - struct.
type RowsInfoToFile struct {
	DatabaseName  string
	DatabaseTable string
	Content       map[int64][]string
}

//RowsInfoToJSON - struct.
type RowsInfoToJSON struct {
	DatabaseName  string
	DatabaseTable string
	Content       map[int64][]string
}

//DBWrite - interface. after setting the content provides ways to save this content
type DBWrite interface {
	SetContent(theContent map[int64][]string)
	Write(targetPath string) error
}

//Write - method. signature satisfies DBWrite interface. takes content from receiver and writes it to CSV file
func (receiver *RowsInfoToFile) Write(targetPath string) error {
	targetFile, err := os.Create(targetPath)
	defer func() {
		if targetFile != nil {
			targetFile.Close()
		}
	}()

	if err != nil {
		log.Fatal("os.Create:", err)
	}

	csvWriter := csv.NewWriter(targetFile)
	defer csvWriter.Flush()

	for _, aValue := range receiver.Content {
		err := csvWriter.Write(aValue)
		if err != nil {
			log.Fatal("csvWriter.Write:", err)
		}
	}
	return nil
}

//SetContent - signature satisfies DBWrite interface. setter for content that should be exported by Write method
func (receiver *RowsInfoToFile) SetContent(theContent map[int64][]string) {
	receiver.Content = theContent
}

//Write - method. signature satisfies DBWrite interface. takes content from receiver and writes it to file in JSON format
func (receiver *RowsInfoToJSON) Write(targetPath string) error {
	targetFile, err := os.Create(targetPath)
	defer func() {
		if targetFile != nil {
			targetFile.Close()
		}
	}()

	if err != nil {
		log.Fatal("os.Create:", err)
	}

	columnNames := receiver.Content[0]
	log.Println(columnNames)

	jsonStruct := make(map[string]interface{}, 0)
	jsonStruct["Database"] = receiver.DatabaseName
	jsonStruct["Table"] = receiver.DatabaseTable

	rawJSONData := make([]map[string]string, 0)

	colCount := len(receiver.Content[0])

	for i, aValue := range receiver.Content {
		if i > 0 {
			rowJSON := make(map[string]string, 0)

			for j := 0; j < colCount; j++ {
				rowJSON[receiver.Content[0][j]] = aValue[j]
			}
			rawJSONData = append(rawJSONData, rowJSON)
		}
	}

	jsonStruct["Data"] = rawJSONData

	jsonRAW, err := json.Marshal(jsonStruct)
	if err != nil {
		log.Fatal("json.Marshal:", err)
	}

	err = ioutil.WriteFile(targetPath, jsonRAW, 0777)
	if err != nil {
		log.Fatal("ioutil.WriteFile:", err)
	}

	return nil
}

//SetContent - signature satisfies DBWrite interface. setter for content that should be exported by Write method
func (receiver *RowsInfoToJSON) SetContent(theContent map[int64][]string) {
	receiver.Content = theContent
}
