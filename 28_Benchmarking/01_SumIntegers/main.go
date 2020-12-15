package main

func main() {

}

func sumLoop(pIntegers []int) int {
	if len(pIntegers) == 0 {
		return 0
	}
	var result int
	for _, v := range pIntegers {
		result = result + v
	}
	return result
}

func sumRecurs(pIntegers []int) int {
	if len(pIntegers) == 0 {
		return 0
	}
	return pIntegers[0] + sumRecurs(pIntegers[1:])
}
