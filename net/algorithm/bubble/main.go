package main

import "fmt"

var array []int

func bubbleSort(array []int) []int {
	var length int
	length = len(array)
	var temp int
	for i := 0;i < length-1 ;i++ {

		for v:= i+1;v < length ;v++ {
			if array[i] > array[v] {
				temp = array[i]
				array[i] = array[v]
				array[v] = temp
			}
		}
	}

	return array
}

func main() {
	array = []int{18, 58, 29, 27, 782, 88, 98, 109, 190, 358, 12, 15, 18, 18, 17, 19, 30, 20, 39, 100}
	new_array := bubbleSort(array)
	fmt.Println(new_array)
}