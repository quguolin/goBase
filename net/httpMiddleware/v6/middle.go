package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type HandleContext func(*Context)

type Context struct {
	context.Context
	Request *http.Request
	Writer http.ResponseWriter
	middle []HandleContext
	index  int
}

type Server struct {
	route  map[string]HandleFunc
	handle []HandleContext
}

func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.middle); c.index++{
		c.middle[c.index](c)
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

func withMiddTime() HandleContext {
	return func(c *Context) {
		t := time.Now()
		defer func() {
			fmt.Println("withMiddTime end time",time.Since(t))
		}()
		fmt.Println("withMiddTime start ",time.Since(t))
		c.Next()
	}
}

func withMiddLog() HandleContext {
	return func(c *Context) {
		//log.Printf("MiddleWare(withMiddLog) Request URL(%s) Method(%v) ", c.Request.URL, c.Request.Method)
		fmt.Println("withMiddLog ")
		c.Next()
	}
}

func (s *Server) createContext(w http.ResponseWriter, req *http.Request, middle []HandleContext) *Context{
	return &Context{
		Request: req,
		Writer:  w,
		middle:  middle,
		index:-1,
	}
}

func (s *Server)routeHandler(path string,h HandleFunc){
	s.route[path] = h
}

func (s *Server) Register(path string, f ...HandleContext) {
	handleNew := make([]HandleContext,0,len(s.handle)+len(f))
	handleNew = append(handleNew,s.handle...)
	handleNew = append(handleNew,f...)
	s.routeHandler(path, func(writer http.ResponseWriter, request *http.Request) {
		s.createContext(writer,request,handleNew).Next()
	})
}

func (s *Server) UseMiddle(hc ...HandleContext){
	s.handle = append(s.handle,hc...)
}

func New() *Server {
	s := &Server{
		route: make(map[string]HandleFunc),
	}
	s.UseMiddle(withMiddLog(),withMiddTime())
	return s
}

func main() {
	r := New()
	r.Register("/hello", func(c *Context) {
		time.Sleep(time.Second)
		fmt.Println("hello sleep 1 second")
		c.Writer.Write([]byte("hello!\r\n"))
	})
	r.Register("/world", func(c *Context) {
		time.Sleep(2 * time.Second)
		fmt.Println("world sleep 2 second")
		_,err := c.Writer.Write([]byte("world\r\n"))
		if err != nil{
			fmt.Println(err)
		}
	})
	http.ListenAndServe(":8080", r)
}
