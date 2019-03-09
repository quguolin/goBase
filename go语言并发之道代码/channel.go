package main

import "fmt"

func main() {
	//c := make(chan int)
	var ch chan int
	if ch == nil {
		fmt.Println("nil")
	}
	//<-ch
	//return
	//close(ch) //close nil channel. panic: close of nil channel
	//ch <- 1 		//fatal error: all goroutines are asleep - deadlock!
	ch = make(chan int)
	//<-ch
	//return
	//if ch == nil {
	//	fmt.Println("nil222222")
	//}
	//ch <- 1 //fatal error: all goroutines are asleep - deadlock!
	//close(ch)
	close(ch) //double close channel。 panic: close of closed channel
	ch <- 1   //send to close channel。 panic: send on closed channel
	//close(ch)
}
