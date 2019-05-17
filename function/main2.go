package main

import "fmt"

func func2(f func(string) string) func() string {
	//result := f("David")
	//fmt.Println(result) // Prints "Hiya, David"
	return func() string {
		return f("func3")
	}
}

func main() {
	anon := func(name string) string {
		return "Hello, " + name
	}
	v := func2(anon)()
	fmt.Println(v)
}
