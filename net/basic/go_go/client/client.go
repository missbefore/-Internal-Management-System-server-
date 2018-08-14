package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error Dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdout)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")

	for  {
		fmt.Println("What to send to the server? 起风了.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " say: " + trimmedInput))
	}

}
