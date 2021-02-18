package main

import (
	"errors"
	"fmt"
)

// should test if integer data is in given sequence
// ex. sequence is 1,2,3 with data being 1,2,3,1,2,3,4.

func main() {
	data := []int{1, 2, 3, 1, 2, 3, 1, 4}
	seq := []int{1, 2, 3}

	fmt.Println(validate(data, seq))
}

func validate(data, seq []int) error {
	if len(data) == 0 || len(seq) == 0 {
		return errors.New("please review passed conditions")
	}

	l := len(seq)

	for i, _ := range data {
		if i%l == 0 {
			fmt.Println("i:", i)
			for ix, pos := range seq {
				if data[i+ix] == pos {
					continue
				}

				fmt.Println("position:", pos, "value not in sequence:", data[i+ix])
				return errors.New("not in sequence")
			}

		}
	}
	return nil
}
