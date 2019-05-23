package main

import (
	"fmt"
	"time"
)

var (
	bucket chan bool
)

const maxCount = 3

//只有拿到bucket的 才允许允许
func work2(i int)  {
	for{
		<-bucket
		fmt.Println(i,time.Now())
	}
}

//并发控制 每秒最多maxCount次
func bucketAllow()  {
	for{
		for i:=0;i< maxCount;i++{
			select {
			case bucket<-true:
			default:
			}
		}
		time.Sleep(time.Second)
	}
}

func main()  {
	bucket = make(chan bool, maxCount)
	go bucketAllow()
	for i:=0;i<100;i++{
		go work2(i)
	}

	time.Sleep(10*time.Second)
}
