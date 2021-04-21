package main

import (
	"strings"
)

func main() {

}

func concaSlice() string {
	x := "1 2 3"
	y := strings.Split(x, " ")

	return strings.Join(y, "")
}

func concaBuilder() string {
	x := "1 2 3"
	y := strings.Split(x, " ")

	var b strings.Builder

	for _, v := range y {
		b.WriteString(v)
	}

	return b.String()
}
