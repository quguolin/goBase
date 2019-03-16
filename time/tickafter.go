package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Tick(time.Second)
	a := time.After(time.Second)
	for {
		select {
		case t := <-t:
			fmt.Println("tick:", t.Unix())
		case t := <-a:
			fmt.Println("after", t.Unix())
		}
	}
}
