package main

import (
	"fmt"
	"sync"
)

func gen(values ...int) (ch chan int) {
	ch = make(chan int)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return
}

func sq(values chan int) (ch chan int) {
	ch = make(chan int)
	go func() {
		defer close(ch)
		for v := range values {
			ch <- v * v
		}
	}()
	return
}

func merge(cs ...chan int) (out chan int) {
	out = make(chan int)
	wg := &sync.WaitGroup{}
	output := func(ch chan int) {
		for v := range ch {
			out <- v
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, v := range cs {
		go output(v)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return
}

func main() {
	//for v := range sq(gen(1, 2, 3)) {
	//	fmt.Println(v)
	//}
	c1 := sq(gen(1, 2, 3))
	c2 := sq(gen(4, 5, 6))
	for v := range merge(c1, c2) {
		fmt.Println(v)
	}
}
