package main

import (
	"fmt"
	"math/rand"
)

func main() {
	f := func() chan int {
		ch := make(chan int)
		go func() {
			defer fmt.Println("rand exit...")
			defer close(ch)
			for {
				ch <- rand.Int()
			}
		}()
		return ch
	}

	for i := 0; i < 3; i++ {
		v := <-f()
		fmt.Println(v)
	}
}
