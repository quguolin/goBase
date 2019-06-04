package main

import (
	"container/list"
	"fmt"
)

const SIZE  = 3

type I interface {

}

type Element struct {
	Key   I
	Value I
}

type Map map[I]*list.Element

type LRU struct {
	cache Map
	link  *list.List
	size  int
}

func New(size int)*LRU  {
	return &LRU{
		cache: make(Map),
		link:  list.New(),
		size:size,
	}
}

func (l *LRU)Set(key,val I)bool  {
	el := Element{
		Key:key,
		Value:val,
	}
	if e,ok := l.cache[key];ok{
		e.Value = el
		l.link.MoveToFront(e)
		return true
	}
	if len(l.cache)<l.size{
		l.cache[key] = l.link.PushFront(el)
		return true
	}
	e := l.link.Back()
	delete(l.cache,e.Value.(Element).Key)
	l.link.Remove(e)
	l.cache[key] = l.link.PushFront(el)
	return true
}

func (l *LRU)Printf()  {
	e := l.link.Back()
	for i:=0;i<l.size;i++{
		var val I
		val = " "
		if e != nil {
			val = e.Value.(Element).Value
			e = e.Prev()
		}
		fmt.Printf("  %v  |", val)
	}
}

func (l *LRU)Get(key I) (I,bool) {
	val,ok := l.cache[key]
	if !ok{
		return nil,false
	}
	l.link.MoveToFront(val)
	return val.Value.(Element),true
}

func main() {
	c := New(4)
	c.Set("A", "A")
	c.Set("B", "B")
	c.Set("C", "C")
	c.Set("D", "D")
	c.Set("E", "E")
	c.Get("D")
	c.Set("F", "F")

	c.Printf()
}

