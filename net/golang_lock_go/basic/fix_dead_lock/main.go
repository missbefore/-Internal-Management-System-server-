package main

import "fmt"

var c = make(chan int ,2)

func main()  {
	go func1()
	go func2()
	<-c
	<-c
	fmt.Println("ok")
}

func func1()  {
	fmt.Println("here1")
	c <- 99
}

func func2()  {
	fmt.Println("here2")
	c <- 1
}
