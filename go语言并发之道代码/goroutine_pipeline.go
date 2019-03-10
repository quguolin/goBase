package main

import "fmt"

func main() {
	//将一组离散的值 转换为 channel上的值
	repeat := func(done chan int, values ...interface{}) chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case c <- v:
					}
				}
			}
			//for {
			//	select {
			//	case <-done:
			//		return
			//	case c <- value:
			//	}
			//}
		}()
		return c
	}
	take := func(done chan int, values chan interface{}, count int) chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for i := 0; i < count; i++ {
				select {
				case <-done:
					return
				//将一个通道的数据 存到另外一个通道里面去
				case c <- <-values:
				}
			}
		}()
		return c
	}

	done := make(chan int)
	defer close(done)

	for v := range take(done, repeat(done, []interface{}{1, 2}...), 10) {
		fmt.Println(v)
	}
}
