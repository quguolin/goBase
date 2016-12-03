package main

import (
	"fmt"
)

func main() {
	elements := map[string]map[string]string{
		"A": map[string]string{
			"name": "aaa",
			"addr": "111",
		},
		"B": map[string]string{
			"name": "bbb",
			"addr": "222",
		},
		"C": map[string]string{
			"name": "ccc",
			"addr": "333",
		},
	}

	for index, value := range elements {
		fmt.Println(index)
		fmt.Println(value)
	}

	test1 := map[int]map[string]string{
		0: map[string]string{
			"name": "aaa",
			"addr": "111",
		},
		1: map[string]string{
			"name": "bbb",
			"addr": "222",
		},
		2: map[string]string{
			"name": "ccc",
			"addr": "333",
		},
	}

	for index, value := range test1 {
		fmt.Println(index)
		fmt.Println(value)
	}

	test2 := map[int]map[string]string{
		0: map[string]string{
			"name": "aaa",
			"addr": "111",
		},
		1: map[string]string{
			"name": "bbb",
			"addr": "222",
		},
		2: map[string]string{
			"name": "ccc",
			"addr": "333",
		},
	}

	for index, value := range test2 {
		fmt.Println(index)
		fmt.Println(value)
	}
}
