package main

import (
	"sync"
	"time"
)

//go run -race race.go
func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	mutex := sync.Mutex{}
	m2 := make(map[string]int)
	for k, v := range m {
		k := k
		v := v
		go func() {
			mutex.Lock()
			m2[k] = v
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)

}
