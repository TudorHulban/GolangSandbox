package main

import (
	"fmt"
)

func main() {
	fmt.Println(factoI(3))
	fmt.Println(factoR(3))
}

func factoI(n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}

func factoR(n int) int {
	if n == 0 {
		return 1
	}

	return n * factoR(n-1)
}
