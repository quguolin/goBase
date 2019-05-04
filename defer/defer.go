package main

import (
	"fmt"
	"time"
)

func example1()  {
	defer fmt.Println("HelloDefer1")
	defer fmt.Println("HelloDefer2")
	fmt.Println("HelloWorld")
}

func example2() int {
	i := 1
	defer func(){
		fmt.Println("Defer i address is ",&i)
		i += 1
	}()

	return func()int{
		fmt.Println("Return i address is ",&i)
		return i
	}()
}

func example3() *int {
	i := 1
	p := &i
	defer func() {
		*p += 1
		fmt.Println("defer:	", p, *p)
	}()

	return func() *int{
		fmt.Println("Return:	", p, *p)
		return p
	}()
}

func example4()  int {
	defer func() {
		fmt.Println("defer ")
	}()
	fmt.Println("example")
	return func() int {
		fmt.Println("return")
		return 0
	}()
}

func example5()  {
	defer func() {
		fmt.Println("example 5")
	}()
	panic("error")
}

func example6()  {
	defer func() {
		fmt.Println("example6")
	}()
	for{
		time.Sleep(time.Second)
		fmt.Println("sleep")
		return
	}
}

func example7()  {
	for i:=0;i<5;i++{
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main()  {
	//example1()
	//i := example2()
	//iaddr := &i
	//fmt.Println("main value is ",i," address is ",iaddr)
	//i2 := example3()
	//i2v := *i2
	//fmt.Println("main 2 address is ",i2,"  is value ",i2v)
	//fmt.Println(*example3())
	//example4()
	//example5()
	//go example6()
	example7()
	time.Sleep(2*time.Second)
}
