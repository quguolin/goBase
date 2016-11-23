package main

import (
	"fmt"
	"log"
	"net/http"
)

func getValue(w http.ResponseWriter, r *http.Request) {
	//第一种方法 先解析 然后直接获取postform 的值
	r.ParseForm()

	switch r.Method {
	case "GET":

	case "POST":
		fmt.Println(r.PostForm)
	}

	//第二种方法 不解析 直接通过postformvalue获取
	//	switch r.Method {
	//	case "GET":

	//	case "POST":
	//		fmt.Println(r.PostFormValue("key"))
	//		fmt.Println(r.PostFormValue("value"))
	//	default:

	//	}
}
func main() {
	http.HandleFunc("/", getValue)
	err := http.ListenAndServe(":8080", nil) //设置监听端口

	if err != nil {
		log.Fatal("listenAndServer:", err)
	}
}
