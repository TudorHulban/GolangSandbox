### Basic sorting
```go
package main

import (
	"fmt"
)

func main() {
	a := []int{2, 7, 1, 1, 3}

	fmt.Println("Final bubble:", bubble(a))
	fmt.Println("Final selection:", selection(a))
	fmt.Println("Final insertion:", insertion1(a))
	fmt.Println("Final free style:", insertion2(a))
}

func bubble(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}

	return a
}

func selection(a []int) []int {
	for i := 0; i < len(a); i++ {
		ixMinim := i

		for j := i + 1; j < len(a); j++ {
			if a[j] < a[ixMinim] {
				a[j], a[ixMinim] = a[ixMinim], a[j]
				ixMinim = j
			}
		}
	}

	return a
}

func insertion1(a []int) []int {
	for i := 1; i < len(a); i++ {
		j := i

		for j > 0 {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}

			j = j - 1
		}
	}

	return a
}

func insertion2(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		currentVal := a[i]
		j := i - 1

		for j >= 0 && a[j] > currentVal {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = currentVal
	}

	return a
}
```
