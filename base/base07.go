package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

func (h *Human) SayHi() {
	fmt.Println("Human SayHi method:", h.name, h.phone)
}

//重写父类方法
func (h *Student) SayHi() {
	fmt.Println("Student SayHi method:", h.name, h.phone)
}
func main() {
	mark := Student{Human{"Mark", 25, "111111"}, "MIT"}
	//	fmt.Println(mark)
	mark.SayHi()
}
