package main

import (
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func valueExact(json, value string) bool {
	jq := gojsonq.New().JSONString(json)
	res := jq.From("rows").Where("departname", "=", "HR").Count()

	fmt.Println("res: ", res)
	return false
}
