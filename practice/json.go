package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	example1()
}

//https://blog.golang.org/json-and-go

func example3() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal(err)
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func example2() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	var str string
	for _, value := range b {
		str = str + string(value)
	}

	fmt.Println(str)
}
func example1() {
	type road struct {
		Name   string
		Number int
	}

	roads := []road{
		{"AAA", 111},
		{"BBB", 222},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}
	var str string
	for _, value := range b {
		str = str + string(value)
	}

	fmt.Println(str)
}
