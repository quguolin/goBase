package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	doHttpGet := func(done chan interface{}, urls ...string) chan *Result {
		results := make(chan *Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				response, err := http.Get(url)
				result := &Result{
					Error:    err,
					Response: response,
				}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	for response := range doHttpGet(done, []string{"https://www.baidu.com", "https://www.google.com.hk", "https://badhost"}...) {
		if response.Error != nil {
			fmt.Println(response.Error)
		} else {
			fmt.Println(response.Response.Status)
		}
	}
}
