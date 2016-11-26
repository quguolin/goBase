package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)

	go func() {
		fmt.Println("GO Runtime and Channel!")
		c <- true
	}()

	<-c //程序在这里被阻塞

	fmt.Println("main function end--------")
}
