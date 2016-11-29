package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 0
	for i := 0; i < 10000000; i++ {
		a += 1
	}

	fmt.Println(index, a)
	c <- true
}
