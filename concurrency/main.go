package main

import (
	"fmt"
	"time"
)


func work(i int)  {
	for{
		fmt.Println(i,time.Now())
	}
}

func main()  {
	for i:=0;i< 10;i++{
		go work(i)
	}

	time.Sleep(time.Second)
}
