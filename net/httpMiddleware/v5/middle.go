package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handle func(http.ResponseWriter, *http.Request)

type HandlerFunc func(*Context)

type Context struct {
	context.Context
	Request *http.Request
	Writer http.ResponseWriter
}

type Server struct {
	route  map[string]Handle
	middle []Handle
	index  int
}

func (r *Server) Next(w http.ResponseWriter, req *http.Request) {
	for ; r.index < len(r.middle); r.index++ {
		r.middle[r.index](w, req)
	}
}

func (r *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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

func (r *Server) createContext(w http.ResponseWriter, req *http.Request) *Context{
	return &Context{
		Request:req,
		Writer:w,
	}
}

func (r *Server) Register(route string, f HandlerFunc) {
	r.route[route] = withMiddLog(withMiddTime(func(writer http.ResponseWriter, request *http.Request) {
		f(r.createContext(writer,request))
	}))
}

func New() *Server {
	return &Server{
		route: make(map[string]Handle),
	}
}

func main() {
	r := New()
	r.Register("/bench", func(c *Context) {
		time.Sleep(time.Second)
		fmt.Println("bench sleep 1 second")
		c.Writer.Write([]byte("hello!\r\n"))
	})
	r.Register("/hello", func(c *Context) {
		time.Sleep(2 * time.Second)
		fmt.Println("hello sleep 2 second")
		_,err := c.Writer.Write([]byte("world\r\n"))
		if err != nil{
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":8080", r)
}
