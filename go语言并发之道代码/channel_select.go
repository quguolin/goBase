package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("blocking....")
	select {
	case <-c:
		fmt.Print("unblocked %v later", time.Since(start))
	}

	fmt.Println("done......")
}
