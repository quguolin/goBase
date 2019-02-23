package main

import (
	"context"
	"fmt"
	"time"
)

func work2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value("name"), "done...")
			return
		default:
			fmt.Println(ctx.Value("name"), "working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c1 := context.WithValue(ctx, "name", "1")
	go work2(c1)
	c2 := context.WithValue(ctx, "name", "2")
	go work2(c2)
	c3 := context.WithValue(ctx, "name", "3")
	go work2(c3)
	time.Sleep(time.Second)
	fmt.Println("stop...")
	cancel()
	time.Sleep(1 * time.Second)
}
