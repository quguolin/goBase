package main

import (
	"fmt"
	"time"
)

func main() {
	//start := time.Now()
	//var c1, c2 <-chan int
	//select {
	//case <-c1:
	//case <-c2:
	//default:
	//	fmt.Printf("after %v", time.Since(start))
	//}

	var c <-chan int
	select {
	case <-c:
	case <-time.After(time.Second):
		fmt.Println("timeout ")
	}
}
