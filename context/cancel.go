package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context is done")
				return
			default:
				fmt.Println("context is runing")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	fmt.Println("context is start...")
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(time.Second)
	fmt.Println("context is end...")
}
