package main

import "fmt"

func main()  {
		/*for i:=0; i <= 5; i++ {
		LABEL1:
			for j:=0; j <= 5 ; j++ {
				if j == 4 {
					continue LABEL1
				}
				fmt.Printf("i: %d, and j: %d\n", i, j)
			}
		}*/
	a := 1
	b := 9
TARGET: // compile error

	goto TARGET
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
