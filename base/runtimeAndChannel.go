package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)

	go func() {
		fmt.Println("GO Runtime and Channel!")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("main function end--------")
}
