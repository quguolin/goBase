package main

import (
	"fmt"
)

func test1()  {
	vals := make([]int, 5)
	fmt.Println(vals)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}

func test2()  {
	vals := make([]int, 0,5)
	fmt.Println(vals)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}

type Student struct {
	Name string
}

func test3()  {
	vals := make([]*Student,1)
	fmt.Println(vals)
}

func test4()  {
	vals := make([]int, 5)
	fmt.Println("Capacity was:", cap(vals))
	fmt.Println("Length was:", len(vals))
	for i := 0; i < 6; i++ {
		vals = append(vals, i)
		fmt.Println("Capacity is now:", cap(vals))
		fmt.Println("Length is now:", len(vals))
	}

	fmt.Println(vals)
}

type Context struct {
	index int
}


func main()  {
	//test1()
	//test2()
	//test3()
	//test4()
}
