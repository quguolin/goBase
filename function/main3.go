package main

import (
	"fmt"
	"time"
)

//func func3(f func(string) string) func(string) string {
//	return func(s string) string {
//		t := time.Now()
//		defer func() {
//			consume := time.Since(t)
//			fmt.Println("time consume ", consume)
//		}()
//		return f(s)
//	}
//}

func tmp(name string) {
	fmt.Println("hello ", name)
	time.Sleep(time.Second)
}

func calTime(f func(name string)) func(string) {
	t := time.Now()
	return func(n string) {
		defer func() {
			fmt.Println("time spend is ", time.Since(t))
		}()
		f(n)
	}
}

func main() {
	//s := tmp("world")
	s := calTime(tmp)
	s("world")
	//s := func3(func(name string) string {
	//	time.Sleep(time.Second)
	//	return "Hello, " + name
	//})("hello")
}
