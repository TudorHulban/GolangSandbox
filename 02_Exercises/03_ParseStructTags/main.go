package main

import (
	"log"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type Column struct {
	ColumnName string
	IsNullable bool
	DataType   string
	Unique     bool
	Index      bool
	Default    string // default value for column
}

type User struct {
	name string `hera:"name-field1"`
	age  int    `hera:"name-field2"`
}

func getTableName(model interface{}) string {
	if t := reflect.TypeOf(model); t.Kind() == reflect.Ptr {
		return inflection.Plural(strcase.ToSnake(t.Elem().Name()))
	}

	return inflection.Plural(strcase.ToSnake(reflect.TypeOf(model).Name()))
}

func getFields(model interface{}) []*Column {
	var (
		setFieldType string // takes field type from tag instead
		allowedTypes = []string{"string", "int", "int64", "time.Time", "float64", "bool"}
		result       []*Column
	)

	val := reflect.ValueOf(model).Elem()
	for i := 0; i < val.NumField(); i++ {
		column := &Column{
			ColumnName: strcase.ToSnake(val.Type().Field(i).Name),
		}

		for _, t := range allowedTypes {
			allowedTypes = append(allowedTypes, "*"+t)
			allowedTypes = append(allowedTypes, "[]"+t)
			allowedTypes = append(allowedTypes, "[]*"+t)
		}

		allowedTypes = append(allowedTypes, "map[string]interface")
		allowedTypes = append(allowedTypes, "interface")

		fieldType := val.Type().Field(i).Type.String()
		tag := val.Type().Field(i).Tag.Get("hera")
		ignoreField := false

		if len(tag) > 0 {
			tags := strings.Split(tag, ",")
			for _, tagS := range tags {
				s := strings.ToLower(strings.TrimSpace(tagS))
				if s == "-" {
					ignoreField = true
				}
				if s == "unique" {
					column.Unique = true
				}
				if s == "index" {
					column.Index = true
				}

				ss := strings.Split(s, "type:")
				if len(ss) > 1 {
					column.DataType = ss[1]
					setFieldType = ss[1]
				}
			}
		}
		log.Println("Tag: ", tag)

		accept := false
		for _, at := range allowedTypes {
			if at == fieldType {
				accept = true
				break
			}
		}

		if ignoreField || !accept {
			continue
		}

		if fieldType == "string" {
			column.DataType = "text"
			column.IsNullable = false
			column.Default = "''"
		}
		if fieldType == "*string" {
			column.DataType = "text"
			column.IsNullable = true
		}
		if fieldType == "[]string" || fieldType == "[]*string" {
			column.DataType = "text[]"
		}
		if fieldType == "int64" || fieldType == "int" {
			column.DataType = "bigint"
			column.IsNullable = false
		}
		if fieldType == "*int64" {
			column.DataType = "bigint"
			column.IsNullable = true
		}
		if fieldType == "[]int64" || fieldType == "[]*int64" {
			column.DataType = "integer[]"
			column.IsNullable = true
		}
		if fieldType == "*time.Time" {
			column.DataType = "timestamptz"
			column.IsNullable = true
		}
		if fieldType == "time.Time" {
			column.IsNullable = false
			column.DataType = "timestamptz"
			column.Default = "NOW()"
		}
		if fieldType == "float64" {
			column.DataType = "numeric"
			column.IsNullable = false
			column.Default = "0.00"
		}
		if fieldType == "*float64" {
			column.DataType = "numeric"
			column.IsNullable = true
		}
		if fieldType == "[]float64" || fieldType == "[]*float64" {
			column.DataType = "numeric[]"
			column.IsNullable = true
		}
		if fieldType == "bool" {
			column.DataType = "boolean"
			column.IsNullable = false
			column.Default = "false"
		}
		if fieldType == "*bool" {
			column.DataType = "boolean"
			column.IsNullable = true
		}
		if fieldType == "[]bool" || fieldType == "[]*bool" {
			column.DataType = "boolean[]"
			column.IsNullable = true
		}
		if fieldType == "map[string]interface" || fieldType == "interface" {
			column.DataType = "jsonb"
			column.IsNullable = true
		}
		if len(setFieldType) > 0 {
			column.DataType = setFieldType
		}

		result = append(result, column)
	}
	return result
}

func main() {
	log.Println("TName: ", getTableName(&User{}))

	ptrFields := getFields(&User{})
	for k, v := range ptrFields {
		log.Println(k, *v)
	}
}
