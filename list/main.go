package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")
	v1 := l.Front()
	//l.Remove(v1)
	fmt.Printf("%v\n", v1.Value)
	a1 := l.Back()
	//l.Remove(a1)
	fmt.Printf("%v\n", a1.Value)
	l.PushFront("e")
	v1 = l.Front()
	//l.Remove(v1)
	fmt.Printf("%v\n", v1.Value)
}