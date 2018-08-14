package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	 sendData(ch)
	 getData(ch)

	//time.Sleep(1e9)
}


func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string)  {
	var input string

	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}
