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
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Println("Human SayHi method:", h.phone, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("Human Sing method:", lyrics)
}

//employee 重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Println("Employee SayHi method:", e.name, e.phone)
}

//定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "111111"}, "MIT", 1111}
	paul := Student{Human{"paul", 25, "222222"}, "MIT", 1111}
	sam := Employee{Human{"sam", 25, "333333"}, "MIT", 1111}
	Tom := Employee{Human{"Tom", 25, "444444"}, "MIT", 1111}

	//定义men类型的变量
	var i Men
	i = mike
	i.SayHi()
	i.Sing("sing---------")

	i = Tom
	i.SayHi()
	i.Sing("Tom sing------")

	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
