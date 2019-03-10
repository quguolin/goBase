package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func(c chan int) {
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				fmt.Println(i)
				c <- i
			}
		}
	}(c2)
	go func(c chan int) {
		for i := 1; i < 100; i++ {
			if i%2 == 1 {
				//fmt.Println(i)
				c <- i
			}
		}
	}(c1)
	go func(c1 chan int, c2 chan int) {
		//var c int
		for i := 1; i <= 100; i++ {
			if i%2 == 1 {
				//c = <-c1
				<-c1
				//fmt.Println(<-c1)
			} else {
				//c = <-c2
				<-c2
				//fmt.Println(<-c2)
			}
			//fmt.Println(c)
		}
	}(c1, c2)

	time.Sleep(time.Second)
	close(c1)
	close(c2)
}
