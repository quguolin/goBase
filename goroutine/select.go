package main

import (
	"fmt"
	"time"
)

func work2(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("watching is out...")
			return
		default:
			fmt.Println("watching is working...")
			time.Sleep(time.Second)
		}
	}
}

func stop(ch chan bool) {
	fmt.Println("stop watching")
	ch <- true
}

func main() {
	ch := make(chan bool)
	go work2(ch)
	time.Sleep(3 * time.Second)
	fmt.Println("")
	go stop(ch)
	for {
		time.Sleep(time.Second)
		fmt.Println("main...")
	}
}
