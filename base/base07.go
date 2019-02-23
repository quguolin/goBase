package main

import (
	"fmt"
)

type Human2 struct {
	name  string
	age   int
	phone string
}

type Student2 struct {
	Human2 //匿名字段
	school string
}

type Employee2 struct {
	Human2  //匿名字段
	company string
}

func (h *Human2) SayHi() {
	fmt.Println("Human SayHi method:", h.name, h.phone)
}

//重写父类方法
func (h *Student2) SayHi() {
	fmt.Println("Student SayHi method:", h.name, h.age)
}
func main() {
	mark := Student2{Human2{"Mark", 25, "123456"}, "MIT"}
	//	fmt.Println(mark)
	mark.SayHi()
}
