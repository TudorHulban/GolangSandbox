package series

import (
	"log"
	"testing"
)

/*
func TestChopCSV01(t *testing.T) {
	p := []int{0, 1, 2, 3}
	d, err := NewCSV(11, "sample.csv", p, []int{}, "./spools/1.csv", "./spools/1.xls", 2)
	if err != nil {
		t.Error(err)
	}
	d.Auto()
	if len(d.errorList) == 0 {
		log.Println("error:", d.errorList)
		t.Error(d.errorList[0])
	}
}


func TestChopStringBase(t *testing.T) {
	p := []int{0, 1}
	d, err := NewCSV(12, "sample.csv", p, []int{}, "./spools/2.csv", "./spools/2.xls", 1)
	if err != nil {
		t.Error(err)
	}
	d.Auto()
	if len(d.errorList) == 0 {
		log.Println("error:", d.errorList)
		t.Error(d.errorList[0])
	}
	log.Println("chopped data: ", d.ChoppedData)
}
*/

func TestChopStream(t *testing.T) {
	p := []int{1, 2, 3}
	s := make([][]interface{}, 0)
	s = append(s, []interface{}{"a", 1, 3, 23})
	s = append(s, []interface{}{"b", 2, 5, 25})
	s = append(s, []interface{}{"c", 3, 7, 27})
	s = append(s, []interface{}{"d", 4, 11, 31})
	s = append(s, []interface{}{"e", 5, 13, 33})
	s = append(s, []interface{}{"f", 6, 16, 36})

	d, err := NewStream(100, s, p, []int{1}, "./spools/3_chop_stream.csv", "./spools/3.xls", 0, 0.18)
	if err != nil {
		t.Error(err)
	}
	d.Auto()
	if len(d.errorList) != 0 {
		log.Println("error:", d.errorList)
		t.Error(d.errorList[0])
		return
	}
	log.Println(d.ChoppedData)
}

/*
func TestCombinerStream(t *testing.T) {
	p := []int{1, 2}
	s := make([][]interface{}, 0)
	s = append(s, []interface{}{"a", 1, 3})
	s = append(s, []interface{}{"b", 2, 5})

	d, err := NewStream(200, s, p, []int{}, "./spools/4_combiner_stream.csv", "./spools/4.xls", 1)
	if err != nil {
		t.Error(err)
	}
	d.Auto()
	if len(d.errorList) != 0 {
		log.Println("error:", d.errorList)
		t.Error(d.errorList[0])
		return
	}
	log.Println(d.ChoppedData)
}
*/
