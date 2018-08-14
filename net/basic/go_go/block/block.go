package main

import (
	"fmt"
	"time"
)

func f1(int chan int)  {
	fmt.Println(<-int)

}

func main()  {
	out := make(chan int)
	go func(ch chan int) {
		ch <- 5
	}(out)
	go f1(out)
	time.Sleep(1e9)
}
