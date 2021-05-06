package main

import (
	"log"
	"reflect"
)

type State struct {
	field1 string  `tformname:"Field One" tformvalues:"a|b|c"`
	field2 int     `tformname:"Field Two" tisvalidrange:"0|10"`
	field3 float64 `tformname:"Field Three" tisvalidrange:"0.1|0.7"`
}

func main() {
	s := State{
		field1: "field1",
		field2: 2,
		field3: .21,
	}

	analyze(s)
}

func analyze(s State) {
	root := reflect.TypeOf(&s).Elem()

	for i := 0; i < root.NumField(); i++ {
		field := root.FieldByIndex([]int{i})
		log.Println(field.Tag.Get("tformname"))
	}
}
