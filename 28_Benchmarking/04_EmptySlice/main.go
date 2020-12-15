package main

// based on benchmark the performance of nil is very slightly better.

func main() {

}

func byLength(slice []string) bool {
	return len(slice) == 0
}

func byNil(slice []string) bool {
	return slice == nil
}
