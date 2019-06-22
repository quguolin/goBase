package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handle func(http.ResponseWriter, *http.Request)

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
			log.Println("time spend is ", time.Since(t))
		}()
		h(writer, request)
	}
}

func withMiddLog(h Handle) Handle {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Request URL(%s) Method(%v) ", request.URL, request.Method)
		h(writer, request)
	}
}

func (r *Router) Register(route string, f Handle) {
	r.route[route] = withMiddLog(withMiddTime(f))
}

func New() *Router {
	return &Router{
		route: make(map[string]Handle),
	}
}

func main() {
	r := New()
	r.Register("/bench", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Second)
		fmt.Println("bench sleep 1 second")
	})
	r.Register("/hello", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Println("hello sleep 2 second")
	})
	http.ListenAndServe(":8080", r)
}
