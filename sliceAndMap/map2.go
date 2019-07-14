package main

import (
	"fmt"
	"sync"
)

var counter = struct {
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

func main() {
	go func() {
		for {
			counter.RLock()
			n := counter.m["key"]
			counter.RUnlock()
			fmt.Println("key:", n)
		}
	}()
	go func() {
		for {
			counter.Lock()
			counter.m["key"]++
			counter.Unlock()
		}
	}()
	select {}
}
