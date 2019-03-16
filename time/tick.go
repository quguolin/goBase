package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range time.Tick(time.Millisecond) {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
