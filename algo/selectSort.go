package main

import (
	"fmt"
)

func selectSort(item []int) {
	for i := 0; i < len(item); i++ {
		min := i
		for j := i + 1; j < len(item); j++ {
			if item[j] < item[min] {
				min = j
			}
		}

		if min != i {
			tmp := item[i]
			item[i] = item[min]
			item[min] = tmp
		}
	}
}

func main() {
	array_item := []int{5, 4, 3, 2, 1}
	selectSort(array_item)

	fmt.Println(array_item)
}
