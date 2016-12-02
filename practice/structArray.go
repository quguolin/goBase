package main

import (
	"fmt"
)

type example struct {
	text []string
}

func main() {
	arr := []example{
		example{
			[]string{"a", "b", "c"},
		},
	}

	fmt.Println(arr)
}
