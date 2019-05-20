package main

import (
	"fmt"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request)

type Router struct {
	route map[string]Handle
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.String()
	if v, ok := r.route[path]; ok {
		v(w, req)
		return
	}
	fmt.Println("error")
}

func (r *Router) Register(route string, f Handle) {
	r.route[route] = f
}

func New() *Router {
	return &Router{
		route: make(map[string]Handle),
	}
}

func main() {
	r := New()
	r.Register("/bench", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("bench")
	})
	r.Register("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello")
	})
	http.ListenAndServe(":8080", r)
}
