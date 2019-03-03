package main

import (
	"fmt"
	"sync"
)

func example1() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

//todo 闭包在使用变量V之前 字符串的迭代已经结束
//todo go语言会将对变量V的引用任然保留 由内存转移到堆 以便goroutine可以继续访问它
func example2() {
	var wg sync.WaitGroup

	wg.Add(3)
	for _, v := range []string{"hello", "greetings", "good day"} {
		fmt.Println(v + "---111")
		go func(v string) {
			defer wg.Done()
			fmt.Println(v)
		}(v)
	}
	wg.Wait()
}

func main() {
	example2()
	//example1()
}
