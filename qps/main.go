package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var total int64

func calQPS() {
	var lastTotal int64
	var lastTime time.Time
	for {
		time.Sleep(time.Second)
		t := atomic.LoadInt64(&total)
		now := time.Now()
		gap := now.Sub(lastTime)
		change := (t - lastTotal) / (gap.Nanoseconds() / 1e9)
		s := fmt.Sprintf("qps is %d/秒", change)
		fmt.Println(s)
		lastTotal = t
		lastTime = now
	}
}

func timer() {
	timer1 := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-timer1.C:
			total++
		}
	}
}


func main() {
	go calQPS()
	go timer()
	for {
		time.Sleep(time.Second)
	}
}
