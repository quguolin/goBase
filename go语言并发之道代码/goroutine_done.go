package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings chan string) chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exit...")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})

	strings := make(chan string)

	terminated := doWork(done, strings)

	go func(strings chan string) {
		for _, v := range []string{"hello", "world"} {
			strings <- v
		}
	}(strings)

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("start cancel doWork")
		close(done)
	}()

	//一直等到 close channel之后 才会 真正的结束 main
	<-terminated
	fmt.Println("exit...")
}
