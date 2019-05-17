package main

import (
	"fmt"
	"time"
)

func sleep() {
	//fmt.Println("hello ", s)
	time.Sleep(time.Second)
}

func cal(f func()) func() {
	return func() {
		t := time.Now()
		defer func() {
			fmt.Println("time spend is ", time.Since(t))
		}()
		f()
	}
}

func main() {
	cal(sleep)()
}
