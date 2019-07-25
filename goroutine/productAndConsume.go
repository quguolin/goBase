package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	product := func() (ch chan int) {
		defer wg.Done()
		ch = make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 10; i++ {
				ch <- i
			}
		}()
		return
	}
	go func(ch chan int) {
		defer wg.Done()
		for {
			msg, ok := <-ch
			if !ok {
				return
			}
			fmt.Println(msg)
		}
	}(product())
	wg.Wait()
}
