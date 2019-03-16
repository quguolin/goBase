package main

import (
	"fmt"
	"net/http"
	"time"
)

type HelloHandle struct {
}

func (h HelloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloHandle")
}

func log2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", log2(HelloHandle{}))
	server.ListenAndServe()
}
