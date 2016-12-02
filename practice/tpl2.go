package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name  string
	Email []string
}

func main() {

	http.HandleFunc("/", temp) //设置访问路由

	err := http.ListenAndServe(":8080", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func temp(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./tpl2.html")
	if err != nil {
		log.Fatal(err)
	}

	//	Person := Person{
	//		Name: "Json",
	//		Email: []string{
	//			"aa@163.com",
	//			"bb@163.com",
	//		},
	//	}

	Person := []string{
		"11111111111",
		"222222222222",
		"3333333333333",
	}
	err = t.Execute(w, Person)
	if err != nil {
		log.Fatal(err)
	}
}
