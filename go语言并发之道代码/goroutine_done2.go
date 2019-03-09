package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	f := func(done chan int) chan int {
		//负责创建goroutine的goroutine 也确保它可以停止goroutine
		ch := make(chan int)
		go func() {
			defer fmt.Println("rand exit...")
			defer close(ch)
			for {
				select {
				case <-done:
					fmt.Println("done...")
					return
				case ch <- rand.Int():
					fmt.Println("----------")
				}
			}
		}()
		return ch
	}

	done := make(chan int)
	for i := 0; i < 3; i++ {
		v := <-f(done)
		fmt.Println(v)
	}
	time.Sleep(3 * time.Second)
	close(done)

	fmt.Println("close")
	time.Sleep(3 * time.Second)
}
