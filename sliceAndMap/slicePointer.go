package main

import (
	"fmt"
	"strconv"
)

func example1() {
	var tmp []*string
	var numberStr string
	for i := 0; i < 10; i++ {
		numberStr = fmt.Sprintf("#%s", strconv.Itoa(i))
		//append的是变量的地址
		tmp = append(tmp, &numberStr)
	}
	//取地址中最后一次更新的值
	for _, n := range tmp {
		fmt.Printf("%s\n", *n)
	}
}

func example2() {
	var tmp []*string
	for i := 0; i < 10; i++ {
		//每次都会生成一个新的地址
		var numberStr string
		numberStr = fmt.Sprintf("#%s", strconv.Itoa(i))
		tmp = append(tmp, &numberStr)
	}
	for _, n := range tmp {
		fmt.Printf("%s\n", *n)
	}
}

func main() {
	example1()
	example2()
}
