package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      //匿名字段 struct
	Skills     //匿名字段 自定义类型
	int        //内置类型作为匿名字段
	speciality string
}

func main() {
	//	//按照顺序初始化
	//	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}

	//	fmt.Println(mark.name)
	//	fmt.Println(mark.age)
	//	fmt.Println(mark.weight)
	//	fmt.Println(mark.speciality)

	//按照字段初始化
	jane := Student{Human: Human{"Mark", 25, 120}, speciality: "Computer Science"}

	fmt.Println(jane.name)
	fmt.Println(jane.age)
	fmt.Println(jane.weight)
	fmt.Println(jane.speciality)

	fmt.Println(jane.Skills)
	fmt.Println(jane.int)

	jane.Skills = []string{"anatomy"}
	jane.int = 3

	fmt.Println(jane)
}
