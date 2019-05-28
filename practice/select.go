package main

import (
	"fmt"
	"time"
)

//
//除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
//除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
//如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
//如果没有 default，那么 代码块会被阻塞，指导有一个 case 通过评估；否则一直阻塞

func main() {
	var c chan int
	c = make(chan int)
	timeout := time.After(5 * time.Second)
	select {
	case s := <-c:
		fmt.Println(s)
	case <-timeout:
		fmt.Println("You talk too much.")
		return
	//default:
	//	fmt.Println("hello")
	}
}
