package main

import "fmt"

func main()  {
	go func1()
	go func2()
	select {
	
	}
}

func func1() {
	fmt.Println("here1")
}

func func2() {
	fmt.Println("here2")
}