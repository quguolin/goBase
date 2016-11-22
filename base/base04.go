package main

import (
	"fmt"
)

type testInt func(int) bool //声明一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}

	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}

	return false
}

func filter(slice []int, f testInt) []int {
	var result []int

	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	odd := filter(slice, isEven)
	fmt.Println(odd)

	even := filter(slice, isOdd)
	fmt.Println(even)
}
