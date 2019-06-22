package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handle func(http.ResponseWriter, *http.Request)

type HandlerFunc func()

type Router struct {
	route  map[string]Handle
	middle []Handle
	index  int
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.String()
	if v, ok := r.route[path]; ok {
		v(w, req)
		return
	}
	fmt.Println("error")
}

func withMiddTime(h Handle) Handle {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		defer func() {
			log.Println("MiddleWare(withMiddTime) spend is ", time.Since(t))
		}()
		h(writer, request)
	}
}

func withMiddLog(h Handle) Handle {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("MiddleWare(withMiddLog) Request URL(%s) Method(%v) ", request.URL, request.Method)
		h(writer, request)
	}
}

func (r *Router) Register(route string, f HandlerFunc) {
	r.route[route] = withMiddLog(withMiddTime(func(writer http.ResponseWriter, request *http.Request) {
		f()
	}))
}

func New() *Router {
	return &Router{
		route: make(map[string]Handle),
	}
}

func main() {
	r := New()
	r.Register("/bench", func() {
		time.Sleep(time.Second)
		fmt.Println("bench sleep 1 second")
	})
	r.Register("/hello", func() {
		time.Sleep(2 * time.Second)
		fmt.Println("hello sleep 2 second")
	})
	http.ListenAndServe(":8080", r)
}
