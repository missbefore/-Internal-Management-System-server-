package main

import "fmt"

var (
	firstname, lastname, s string
	i int
	f float32
	input = "56.12/ 5212 / Go"
	format = "%f / %d / %s"
)

func main()  {
	fmt.Println("请填入你的全名: ")
	fmt.Scanln(&firstname, &lastname)
	fmt.Printf("Hi %s %s!\n", firstname, lastname)
}
