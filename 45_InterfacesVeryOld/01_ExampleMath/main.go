package main

import (
	"fmt"
)

func newMath(opAdd mathadd) math {
	return TExtMath{
		a: opAdd.(TAdd),
	}
}

func main() {
	var a TAdd
	fmt.Println(a.add(1, 2))

	var b TMath
	fmt.Println(b.add(3, 4))
	fmt.Println(b.multiply(3, 4))

	var c TExtMath
	fmt.Println(c.add(3, 4))
	fmt.Println(c.multiply(3, 4))

	d := newMath(a)
	fmt.Println(d.add(5, 4))
	fmt.Println(d.multiply(5, 4))
}
