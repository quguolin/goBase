package serverHttp

import (
	"fmt"
	"net/http"
	"runtime"
)

type Context struct {
	Req      *http.Request
	Writer   http.ResponseWriter
	handlers []HandlerFunc
	index    int8
}
type HandlerFunc func(*Context)

type Handle func(http.ResponseWriter, *http.Request)

type Server struct {
	Handlers []HandlerFunc
	router   *Router
}

type Router struct {
	route map[string]Handle
}

func New() *Server {
	engine := &Server{}
	engine.router = &Router{
		route: make(map[string]Handle),
	}
	return engine
}

func Default() *Server {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

// Adds middlewares to the group, see example code in github.
func (group *Server) Use(middlewares ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middlewares...)
}
func Logger() HandlerFunc {
	return func(context *Context) {
		fmt.Println("Logger")
		context.Next()
	}
}

func Recovery() HandlerFunc {
	return func(context *Context) {
		defer func() {
			if i := recover(); i != nil {
				size := 1024 * 1024
				buf := make([]byte, size)
				rs := runtime.Stack(buf, false)
				if rs > size {
					rs = size
				}
				buf = buf[:rs]
				fmt.Println(string(buf))
			}
		}()
		context.Next()
	}
}

func (c *Context) Next() {
	c.index++
	s := int8(len(c.handlers))
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (group *Server) GET(path string, handlers ...HandlerFunc) {
	group.Handle("GET", path, handlers)
}

func (group *Server) Handle(method, p string, handlers []HandlerFunc) {
	handlers = group.combineHandlers(handlers)
	group.router.Handle(method, p, func(w http.ResponseWriter, req *http.Request) {
		group.createContext(w, req, handlers).Next()
	})
}

func (group *Server) createContext(w http.ResponseWriter, req *http.Request, handlers []HandlerFunc) *Context {
	return &Context{
		Writer:   w,
		Req:      req,
		index:    -1,
		handlers: handlers,
	}
}

func (r *Router) Handle(method, path string, handle Handle) {
	r.route[path] = handle
}

func (group *Server) combineHandlers(handlers []HandlerFunc) []HandlerFunc {
	s := len(group.Handlers) + len(handlers)
	h := make([]HandlerFunc, 0, s)
	h = append(h, group.Handlers...)
	h = append(h, handlers...)
	return h
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (engine *Server) Run(addr string) {
	if err := http.ListenAndServe(addr, engine); err != nil {
		panic(err)
	}
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.String()
	if v, ok := r.route[path]; ok {
		v(w, req)
	} else {
		w.Write([]byte("url not found\r\n"))
	}
}

func main() {
	r := Default()
	r.GET("/user", func(context *Context) {
		context.Writer.Write([]byte("hello!\r\n"))
	})
	r.Run(":8081")
}
