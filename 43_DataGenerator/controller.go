package main

import (
	"encoding/json"
	"errors"
)

type Row struct {
	ColumnNames []string
	Values      []interface{}
}

type GenData struct {
	ColumnNames []string        `json:"columns"`
	Rows        [][]interface{} `json:"rowsdata"`
}

type FieldConfig struct {
	Name         string `json:"name"`
	Type         int    `json:"ftype"`
	Length       int    `json:"flength"`
	PositiveOnly bool   `json:"fpositive"`
	MinValue     int    `json:"min"`
	MaxValue     int    `json:"max"`
}

type GenConfig struct {
	Configuration []FieldConfig `json:"config"`
	NoRows        int64         `json:"norows"`
}

func GetConfig(pConfig string) (*GenConfig, error) {
	instance := new(GenConfig)

	err := json.Unmarshal([]byte(pConfig), &instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func GetData(pConfig *GenConfig) (*GenData, error) {
	instance := new(GenData)
	instance.Rows = make([][]interface{}, pConfig.NoRows)
	bufferColumns := [][]interface{}{}
	fieldTpes := []int{} //really dirty

	for _, v := range pConfig.Configuration {
		instance.ColumnNames = append(instance.ColumnNames, v.Name)
		fieldTpes = append(fieldTpes, v.Type)

		switch v.Type {
		case 0:
			{
				if (v.MinValue > v.MaxValue) && (v.MaxValue == 0) {
					return nil, errors.New("Bad Configuration.")
				}
				n := NewNumbers(v.Length, v.PositiveOnly, v.MinValue, v.MaxValue, int(pConfig.NoRows))
				bufferColumns = append(bufferColumns, *SliceIntToInterface(&n))
			}

		case 1:
			{
				s := NewCharacters(v.Length, int(pConfig.NoRows))
				bufferColumns = append(bufferColumns, *SliceStringToInterface(&s))
			}
		case 2:
			{
				id := NewIDs(int(pConfig.NoRows))
				bufferColumns = append(bufferColumns, *SliceIntToInterface(&id))
			}
		}
	}

	for i := 0; i < int(pConfig.NoRows); i++ {
		instance.Rows[i] = make([]interface{}, len(instance.ColumnNames))

		for k, _ := range instance.ColumnNames {
			instance.Rows[i][k] = bufferColumns[k][i]
		}
	}
	return instance, nil
}

func SliceIntToInterface(pSlice *[]int) *[]interface{} {

	instance := make([]interface{}, len(*pSlice))
	for k, v := range *pSlice {
		instance[k] = v
	}
	return &instance
}

func SliceStringToInterface(pSlice *[]string) *[]interface{} {

	instance := make([]interface{}, len(*pSlice))
	for k, v := range *pSlice {
		instance[k] = v
	}
	return &instance
}
