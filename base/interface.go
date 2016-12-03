package main

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}
func main() {
	t := []int{1, 2, 3, 4, 5, 6}
	s := make([]interface{}, len(t))

	for i, v := range t {
		s[i] = v
	}

	fmt.Println(s)

	names := []string{"aaa", "bbb", "ccc"}

	inters := make([]interface{}, len(names))
	for i, val := range names {
		inters[i] = val
	}
	PrintAll(inters)
}
