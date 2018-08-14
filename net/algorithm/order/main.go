package main

import "fmt"

var (
	length int
	array [9]int
	result int
)

func order(length int) int {

	var number int

	for i := 0; i <= length; i++ {
		if i == 0 || i == 1 {
			number = 1
		} else {
			number = array[i-1] + array[i-2]
		}
		array[i] = number
	}
	fmt.Println(array)
	return array[length-1]
}

func main() {
	length = 8
	result = order(length)
	fmt.Println(result)
}
