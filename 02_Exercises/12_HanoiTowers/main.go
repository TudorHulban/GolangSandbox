package main

import (
	"fmt"
)

var (
	rodFrom    = []int{4, 3, 2, 1}
	rodInterim = []int{}
	rodTo      = []int{}
)

var op int

func main() {
	moveDisks(len(rodFrom), &rodFrom, &rodTo, &rodInterim)
}

func moveDisks(n int, from, to, via *[]int) {
	if n == 0 {
		return
	}

	if n == 1 {
		moveDisk(from, to)
		return
	}

	// default
	if len(*from) == 0 {
		return
	}

	moveDisks(n-1, from, to, via)
	moveDisk(from, to)

	if len(*via) == 0 {
		return
	}

	moveDisks(n-1, via, from, to)
}

func moveDisk(from, to *[]int) {
	op++
	fmt.Println(op, "Before:", "rodFrom: ", rodFrom, "rodTo: ", rodTo)

	if len(*from) == 1 {
		*to = append(*to, (*from)[0])
		*from = []int{}

		fmt.Println(op, "After :", "rodFrom: ", rodFrom, "rodTo: ", rodTo)
		return
	}

	lastEl := (*from)[len(*from)-1]

	// from does not support slicing
	a := *from
	a = a[:len(a)-1]
	*from = a

	*to = append(*to, lastEl)

	fmt.Println(op, "After :", "rodFrom: ", rodFrom, "rodTo: ", rodTo)
}
