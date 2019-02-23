package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("gorouting " + name + " done")
			return
		default:
			fmt.Println("gorouting " + name + "workiing...")
			time.Sleep(time.Second)
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("start...")
	go work(ctx, "first")
	go work(ctx, "second")
	go work(ctx, "third")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
