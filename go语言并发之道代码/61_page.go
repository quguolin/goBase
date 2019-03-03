package main

import (
	"fmt"
	"sync"
)

//mutex 互斥锁来同步访问代码123

func main() {
	var count int
	var lock sync.Mutex
	var wg sync.WaitGroup

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("incrementing:%d\n", count)
	}

	size := 10
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
}
