package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func main()  {
	m = new(sync.RWMutex)

	go read(1)
	go read(2)

	time.Sleep(4*time.Second)
}

func read(i int)  {
	println(i, "read start")

	m.RLock()
	println(i,"reading")
	m.RUnlock()

	time.Sleep(3*time.Second)
	println(i, "read over")
}