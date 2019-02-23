package main

import (
	"fmt"
)

//defer 被插入到return 之前执行
func f4() (result int) {
	result = 0
	defer func() {
		result++
	}()

	return result
}

//defer 被插入到赋值和返回之间
func f5() (r int) {
	t := 5
	defer func() {
		t += 5
	}()

	return t
}

func f6() (r int) {
	r = 1
	defer func(r int) {
		r = r + 5 //这里改的r是传值传进去的r，不会改变要返回的那个r值
	}(r)
	return r
}

func main() {
	result1 := f1()
	fmt.Println(result1)

	result2 := f2()
	fmt.Println(result2)

	result3 := f3()
	fmt.Println(result3)
	//	sum := 0

	//	for index := 1; index <= 100; index++ {
	//		sum += index
	//	}

	//	fmt.Println(sum)

	//	//defer 先进后出
	//	for i := 0; i < 5; i++ {
	//		defer fmt.Println(i)
	//	}

}
