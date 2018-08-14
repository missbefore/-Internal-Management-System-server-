package main

import "fmt"
/*
func func1()  {
	fmt.Printf("In function1 at the top\n")
	defer func2()
	fmt.Printf("In function1 at the bottom!\n")
}

func func2()  {
	fmt.Println("shit")
}*/

func main()  {
	a()
}

func a() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n ", i)
	}
}
