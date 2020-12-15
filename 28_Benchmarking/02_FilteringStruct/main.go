package main

import (
	"log"
	"reflect"
	"strconv"
)

type Task struct {
	ID   int
	Name string
	Area string
}

type Work struct {
	Tasks []Task
}

func main() {
	t1 := Task{ID: 1, Name: "x1", Area: "db"}
	t2 := Task{ID: 2, Name: "x2", Area: "os"}
	t3 := Task{ID: 3, Name: "x3", Area: "db"}

	var w1 Work
	w1.Tasks = []Task{t1, t2, t3}

	log.Println(w1.contains(11))
	log.Println(w1.contains(1))

	w2 := *NewTasks(10)
	log.Println(w2.contains(11))
	log.Println(w2.contains(1))

	log.Println(w2.showFieldValues("Name"))
}

func (w *Work) contains(pID int) int {
	for k, v := range w.Tasks {
		if v.ID == pID {
			return k
		}
	}
	return -1
}

func NewTasks(pHowMany int) *Work {
	instance := new(Work)
	for i := 0; i < pHowMany; i++ {
		var task Task
		task.ID = i
		task.Name = "N" + strconv.Itoa(i)
		task.Area = "A" + strconv.Itoa(i)
		instance.Tasks = append(instance.Tasks, task)
	}
	return instance
}

func (w *Work) showFieldValues(pFieldName string) []interface{} {
	var result []interface{}
	for _, v := range w.Tasks {
		result = append(result, reflect.ValueOf(v).FieldByName(pFieldName).Interface())

	}
	return result
}

func showFieldDirect(pFieldName string, pFrom *Work) []interface{} {
	var result []interface{}
	for _, v := range pFrom.Tasks {
		result = append(result, reflect.ValueOf(v).FieldByName(pFieldName).Interface())
	}
	return result
}

func showFieldName(pFieldName string, pFrom *Work) []interface{} {
	var result []interface{}
	for _, v := range pFrom.Tasks {
		result = append(result, v.Name)
	}
	return result
}

func showFieldAssertion(pFieldName string, pFrom interface{}) []interface{} {
	var result []interface{}
	for _, v := range pFrom.(Work).Tasks {
		result = append(result, reflect.ValueOf(v).FieldByName(pFieldName).Interface())
	}
	return result
}
