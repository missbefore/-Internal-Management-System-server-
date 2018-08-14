package main

import (
	"fmt"
	"time"
)

func quickSort(nums []int, ch chan int, level int, threads int, my_level chan int)  {
	level = level * 2
	go func() {
		my_level <-level
		time.Sleep(3 * time.Second)
	}()
	if len(nums) == 1 {
		ch <- nums[0]
		close(ch)
		return
	}
	if len(nums) == 0 {
		close(ch)
		return
	}

	less := make([]int, 0)
	greater := make([]int , 0)
	left := nums[0]
	nums = nums[1:]

	for _, num_data := range nums {
		switch {
		case num_data <= left:
			less = append(less, num_data)
		case num_data > left:
			greater = append(greater, num_data)
		}
	}

	left_ch := make(chan int, len(less))
	right_ch:= make(chan int, len(greater))

	if level <= threads {
		go quickSort(less, left_ch, level, threads, my_level)
		go quickSort(greater, right_ch, level, threads, my_level)
	} else {
		 quickSort(less, left_ch, level, threads, my_level)
		 quickSort(greater, right_ch, level, threads, my_level)
	}

	for i := range left_ch {
		ch <-i
	}
	ch <- left
	for i := range right_ch {
		ch<-i
	}
	close(ch)

	return
}

func main()  {
	array := []int{18, 58, 29, 27, 782, 88, 98, 109, 190, 358, 12, 15, 18, 18, 17, 19, 30, 20, 39, 100}
	ch := make(chan int)
	my_level := make(chan int)
	go quickSort(array, ch, 0, 0, my_level)

	for a := range my_level {
		fmt.Println(a)
	}

}