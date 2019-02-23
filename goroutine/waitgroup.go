package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name + " done")
}

func main() {
	var wg sync.WaitGroup
	number := 3
	wg.Add(number)
	for i := 0; i < number; i++ {
		go work(&wg, strconv.Itoa(i))
	}
	wg.Wait()
	fmt.Println("done...")
}
