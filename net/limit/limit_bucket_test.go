package limit

import (
	"fmt"
	"testing"
	"time"
)

func TestTake(t *testing.T) {
	done := make(chan int)
	c := Bucket(done, 2, 10)

	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(time.Now().Unix())
		if Take(c) {
			fmt.Println("success")
		} else {
			fmt.Println("fail")
		}
	}
	close(done)
	time.Sleep(time.Second)
}
