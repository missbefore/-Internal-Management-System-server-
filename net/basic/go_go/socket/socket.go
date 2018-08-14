package main

import (
	"net"
	"io"
	"fmt"
)

func main()  {
	var (
		host          = "www.apache.org"
		port          = "8060"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)

	conn, err := net.Dial("tcp", remote)
	io.WriteString(conn, msg)

	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}

	conn.Close()
}
