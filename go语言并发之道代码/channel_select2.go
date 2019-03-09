package main

import "fmt"

func main() {

	c1 := make(chan interface{})
	c2 := make(chan interface{})
	close(c1)
	close(c2)

	//select 会在多个case之间做伪随机选择
	//意味着 在case语句集合中 每一个都有一个 被执行的机会
	//defautl 操作 只有所有的case都不通的时候 才会执行
	var count1, count2, count3 int
	for i := 0; i < 1000; i++ {
		select {
		case <-c1:
			count1++
		case <-c2:
			count2++
		default:
			count3++
		}
	}
	fmt.Println(count1, count2, count3)
}
