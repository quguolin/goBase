package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan bool
	ch = make(chan bool)
	go func() {
		for {
			fmt.Println("send")
			ch <- true
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		timer := time.NewTimer(time.Second * 1)
	label1:
		for {
			if !timer.Stop() {
				select {
				case <-timer.C: //try to drain from the channel
				default:
				}
			}
			timer.Reset(time.Second * 1)
			for {
				select {
				case v := <-ch:
					fmt.Println("receive", v)
				//default:
				//	fmt.Println("default")
				case <-timer.C:
					fmt.Println(time.Now(), ":timer expired")
					goto label1
				}
			}
		}
	}()

	var s string
	fmt.Scanln(&s)
}
