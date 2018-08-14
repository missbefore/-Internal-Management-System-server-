package main

import (
	"fmt"
	"time"
)

func main() {
	go fun1()
	go fun2()
	time.Sleep(1 * time.Minute)
}

func fun1() {
	for  {
		fmt.Println("here1")
		time.Sleep(10 * time.Minute)
	}
}

func fun2()  {
	for {
		fmt.Println("here2")
		time.Sleep(10 * time.Minute)
	}
}
