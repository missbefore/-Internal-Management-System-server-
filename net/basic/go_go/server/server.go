package main

import (
	"fmt"
	"net"
)

func main()  {
	fmt.Println("Starting server...")
	
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}

	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error ", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn)  {
	for  {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Receved data:%v", string(buf[:len]))
	}
}
