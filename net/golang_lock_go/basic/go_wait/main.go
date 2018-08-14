package main

import (
	"sync"
	"net/http"
	"io/ioutil"
	"fmt"
)

var wg sync.WaitGroup
var urls = []string{
	"https://www.baidu.com/",
	"http://www.sohu.com/",
	"http://www.sina.com.cn/",
}


func main()  {

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				// handle error
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("%s\r\n",body)
		}(url)
	}

	wg.Wait()
}
