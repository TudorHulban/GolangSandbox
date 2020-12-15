// Bubble Sort in Golang
package main

import (
	"log"
)

func main() {
	slice := generateSlice(10)

	log.Println("Unsorted slice: ", slice)
	bubblesort(slice)
	log.Println("Sorted slice: ", slice)
}
