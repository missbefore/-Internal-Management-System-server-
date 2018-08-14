package main

import "fmt"

var array []int

func quickSort(left int, right int) {

	if left >= right {
		return
	}

	i := left
	j := right
	temp := array[left]

	for i != j {
		for array[j] >= temp && i < j {
			j--
			fmt.Println("+++++")
			fmt.Println(j)
		}

		for array[i] <= temp && i <j {
			i++
			fmt.Println("-----")
			fmt.Println(i)
		}

		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i], array[left] = temp, array[i]
	quickSort(left, i-1)

	if i != len(array)-1 {
		quickSort(i+1, right)
	}
}

func main() {
	array = []int{18, 58, 29, 27, 782, 88, 98, 109, 190, 358, 12, 15, 18, 18, 17, 19, 30, 20, 39, 100}
	quickSort(0, len(array)-1)
	fmt.Println(array)
}
