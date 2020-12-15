package datageneration

import (
	"encoding/json"
	"errors"
)

/*
type row struct {
	ColumnNames []string
	Values      []interface{}
}
*/

// GenData - structure for data generation. final result of the data generation based on input configuration
type GenData struct {
	ColumnNames []string        `json:"columns"`
	Rows        [][]interface{} `json:"rowsdata"`
}

// fieldConfig - configuration for field. attributes used for the data generated in this field
type fieldConfig struct {
	Name         string `json:"name"`
	Type         int    `json:"ftype"`
	Length       int    `json:"flength"`
	PositiveOnly bool   `json:"fpositive"`
	MinValue     int    `json:"min"`
	MaxValue     int    `json:"max"`
}

// GenConfig - holder for field configuration and number of rows to generate
type GenConfig struct {
	Configuration []fieldConfig `json:"config"`
	NoRows        int64         `json:"norows"`
}

// GetConfig - takes configuration as JSON and maps it to the configuration structure
func GetConfig(pConfigJSON string) (*GenConfig, error) {
	instance := new(GenConfig)

	err := json.Unmarshal([]byte(pConfigJSON), &instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// GetData - takes input configuration structure, returns generated data
func GetData(pConfig *GenConfig) (*GenData, error) {
	instance := new(GenData)
	instance.Rows = make([][]interface{}, pConfig.NoRows)
	bufferColumns := [][]interface{}{}

	for _, fieldConfig := range pConfig.Configuration {
		instance.ColumnNames = append(instance.ColumnNames, fieldConfig.Name)

		switch fieldConfig.Type {
		case 0:
			{
				if (fieldConfig.MinValue > fieldConfig.MaxValue) && (fieldConfig.MaxValue == 0) {
					return nil, errors.New("bad configuration")
				}
				n := newNumbers(fieldConfig.Length, fieldConfig.PositiveOnly, fieldConfig.MinValue, fieldConfig.MaxValue, int(pConfig.NoRows))
				bufferColumns = append(bufferColumns, *sliceIntToInterface(&n))
			}

		case 1:
			{
				s := newCharacters(fieldConfig.Length, int(pConfig.NoRows))
				bufferColumns = append(bufferColumns, *sliceStringToInterface(&s))
			}
		case 2:
			{
				id := newIDs(int(pConfig.NoRows))
				bufferColumns = append(bufferColumns, *sliceIntToInterface(&id))
			}
		}
	}

	for i := 0; i < int(pConfig.NoRows); i++ {
		instance.Rows[i] = make([]interface{}, len(instance.ColumnNames))

		for k := range instance.ColumnNames {
			instance.Rows[i][k] = bufferColumns[k][i]
		}
	}
	return instance, nil
}

func sliceIntToInterface(pSlice *[]int) *[]interface{} {
	instance := make([]interface{}, len(*pSlice))

	for k, v := range *pSlice {
		instance[k] = v
	}
	return &instance
}

func sliceStringToInterface(pSlice *[]string) *[]interface{} {
	instance := make([]interface{}, len(*pSlice))

	for k, v := range *pSlice {
		instance[k] = v
	}
	return &instance
}
