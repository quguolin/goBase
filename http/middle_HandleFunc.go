package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "hello\r\n")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		h(writer, request)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/hello", log(hello))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
