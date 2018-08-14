package main

import (
	"time"
	"fmt"
)

func main()  {
	steam := dump()
	suck(steam)
	time.Sleep(1e9)
}

func dump() chan int {
	ch := make(chan int)
	go func() {
		for i:=0; ;i++  {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int)  {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
