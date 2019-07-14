package main

import (
	"fmt"
	"sync"
)

var myMap *sync.Map

func main() {
	myMap = &sync.Map{}
	myMap.Store("key", 1)
	go func() {
		for {
			if v, ok := myMap.Load("key"); ok {
				fmt.Println("key:", v)
			}
		}
	}()
	go func() {
		for {
			if val, ok := myMap.Load("key"); ok {
				v := val.(int)
				v = v + 1
				myMap.Store("key", v)
			}
		}
	}()
	select {}
}
