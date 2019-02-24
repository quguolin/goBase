package main

import (
	"fmt"
	"io"
	"net/http"
)

type Handle func(http.ResponseWriter, *http.Request)

//存储路由对应的handle
type Route struct {
	route map[string]Handle
}

//路由调用回调函数
func (r *Route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.String()
	fmt.Println(path)
	if v, ok := r.route[path]; ok {
		v(w, req)
	}
}

func New() *Route {
	return &Route{
		make(map[string]Handle),
	}
}

func (r *Route) AddRoute(path string, handler Handle) {
	r.route[path] = handler
}

func (r *Route) Run(addr string) {
	http.ListenAndServe(addr, r)
}

func main() {
	r := New()
	r.AddRoute("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello")
		io.WriteString(writer, "hello\r\n")
	})
	r.Run(":8080")
}
