package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	var workCount int
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	for {
		select {
		case <-done:
			fmt.Println(workCount)
			return
		default:
			workCount++
			time.Sleep(time.Second)
		}
	}

}
