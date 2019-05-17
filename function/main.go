package main

import "fmt"

func func1(f func(string) string) {
	result := f("func1")
	fmt.Println(result) // Prints "Hiya, David"
}

func main() {
	anon := func(name string) string {
		return "Hello, " + name
	}
	func1(anon)
}
