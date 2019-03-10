package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	doHttpGet := func(done chan interface{}, urls ...string) chan *http.Response {
		response := make(chan *http.Response)
		go func() {
			fmt.Println("start...")
			defer fmt.Println("close")
			defer close(response)
			for _, url := range urls {
				r, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case response <- r:
				case <-done:
					fmt.Println("done...")
					return
				}
			}
		}()
		return response
	}

	done := make(chan interface{})
	//for v := range doHttpGet(done, []string{"https://www.qq.com", "http://badHost"}...) {
	//	fmt.Println(v.Status)
	//}
	doHttpGet(done, []string{"https://www.qq.com", "http://badHost"}...)
	fmt.Println("start close")
	close(done)
	time.Sleep(2 * time.Second)
}
