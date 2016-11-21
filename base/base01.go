package main

import (
	"fmt"
)

func main() {
	//map 第一种方式
	var m1 map[string]int
	m1 = make(map[string]int)

	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3

	fmt.Println(m1)

	//map 第二种方式
	m2 := make(map[string]int)
	m2["one"] = 1
	m2["two"] = 1
	m2["three"] = 1

	fmt.Println(m2)

	//map 第三种方式
	m3 := map[string]float32{
		"C":      5,
		"Go":     4.5,
		"Python": 4.5,
		"C++":    2,
	}

	fmt.Println(m3)

	//删除map
	delete(m3, "C")
	fmt.Println(m3)
}
