package main

import "fmt"
func main()  {
	var vals []int
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
		fmt.Println("The length of our slice is:", len(vals))
		fmt.Println("The capacity of our slice is:", cap(vals))
	}

	// Add a new item to our array
	vals = append(vals, 123)
	fmt.Println("The length of our slice is:", len(vals))
	fmt.Println("The capacity of our slice is:", cap(vals))

	// Accessing items is the same as an array
	fmt.Println(vals[5])
	fmt.Println(vals[2])
}
