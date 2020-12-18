package main

import (
	"log"
	"reflect"
)

type User struct {
	Name string `json:name-field`
	Age  int
}

func main() {
	user1 := &User{
		Name: "x",
		Age:  20,
	}

	// Reflect on Type
	root1 := reflect.TypeOf(user1).Elem()

	f1 := root1.FieldByIndex([]int{0})
	log.Println(f1.Name, f1.Tag, f1.Type, f1.Index)
	log.Println("Number fields: ", root1.NumField())

	// Reflect on Value
	root2 := reflect.ValueOf(user1).Elem()
	log.Println("Fields: ", root2)
	log.Println("Field value: ", root2.FieldByIndex([]int{0}))

	f5 := reflect.ValueOf(user1).Elem()
	log.Println(f5)
	log.Println(f5.CanSet())

	f5.Field(0).SetString("y")
	f5.Field(1).SetInt(1)
	log.Println(f5, *user1)

	f6 := reflect.New(reflect.TypeOf(user1))
	log.Println("f6: ", f6.CanSet())
}

/*
Result

2009/11/10 23:00:00 Name json:name-field string [0]
2009/11/10 23:00:00 Number fields:  2
2009/11/10 23:00:00 Fields:  {x 20}
2009/11/10 23:00:00 Field value:  x
2009/11/10 23:00:00 {x 20}
2009/11/10 23:00:00 true
2009/11/10 23:00:00 {y 1} {y 1}
2009/11/10 23:00:00 f6:  false
*/
