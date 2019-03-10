package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeat := func(done chan interface{}, fn func() interface{}) chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for {
				select {
				case <-done:
					return
				case c <- fn():
				}
			}
		}()
		return c
	}

	take := func(done chan interface{}, values chan interface{}, count int) chan interface{} {
		ch := make(chan interface{})
		go func() {
			defer close(ch)
			for i := 0; i < count; i++ {
				select {
				case <-done:
					return
				case ch <- <-values:
				}
			}
		}()
		return ch
	}
	done := make(chan interface{})
	rand := func() interface{} {
		return rand.Int()
	}
	for v := range take(done, repeat(done, rand), 10) {
		fmt.Println(v)
	}

	close(done)
	//for v := range repeat(done, rand) {
	//	fmt.Println(v)
	//}
}
