package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func timeprint(p string) {
	time.Sleep(time.Second)
	fmt.Printf("%s time is %d\r\n", p, time.Now().Unix())
}

func a() {
	lock.Lock()
	timeprint("a")
	lock.Unlock()
}

func b() {
	lock.Lock()
	timeprint("b")
	lock.Unlock()
}

func all() {
	lock.Lock()
	timeprint("a")
	timeprint("b")
	lock.Unlock()
}

func main() {
	a()
	b()
	all()
}
