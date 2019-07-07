package limit

import (
	"fmt"
	"time"
)

func Bucket(done chan int, limit, rate int) chan struct{} {
	bc := make(chan struct{}, limit)
	ticker := time.Tick(time.Duration(1) * time.Second)
	//先把桶里面存满令牌
	for i := 0; i < limit; i++ {
		bc <- struct{}{}
	}
	fmt.Println(time.Now().Unix())
	//使用定时器 每秒钟 向channel中存rate个结构体
	go func() {
		defer close(bc)
		select {
		case <-done:
			return
		default:
		}
		for {
			<-ticker
			fmt.Println(time.Now().Unix())
			for i := 0; i < rate; i++ {
				bc <- struct{}{}
			}
		}
	}()
	return bc
}

func Take(c chan struct{}) bool {
	select {
	case <-c:
		return true
	default:
		return false
	}
}
