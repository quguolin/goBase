## golang http 中间件
[源码链接](https://github.com/quguolin/goBase/tree/master/net/httpMiddleware)

> golang的http中间件的实现 首先实现一个http的handler接口

```golang
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

type Router struct {
	route map[string]Handle
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}
```

###  通过函数包裹的方式实现

####中间件v1.0


>1.通过匿名函数 将handler包裹起来 然后在 调用传进来的handler。在执行传进来的参数之前
就可以做到记录日志 等一些中间件的功能

>2.如果有多个中间件 那么就多个函数 一层一层包裹

```golang
func withMiddle(h Handle) Handle {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		defer func() {
			log.Println("time spend is ", time.Since(t))
		}()
		h(writer, request)
	}
}

func (r *Router) Register1(route string, f Handle) {
	r.route[route] = withMiddle(f)
}

func (r *Router) Register2(route string, f Handle) {
	r.route[route] = withMiddLog(withMiddTime(f))
}

Register("/bench", func(writer http.ResponseWriter, request *http.Request) {
    time.Sleep(time.Second)
    fmt.Println("bench sleep 1 second")
})
```


####中间件v1.1

>注册的时候 可以更加简化一些 通过匿名函数的方式 当然这种方式没有传递参数
只是作为演示用的

```golang
func (r *Router) Register(route string, f HandlerFunc) {
	r.route[route] = withMiddLog(withMiddTime(func(writer http.ResponseWriter, request *http.Request) {
		f()
	}))
}
```


####中间件v2.0

>针对中间件v1.1中的没法传递 http中的读写参数的问题 可以封装一个context
将http的读写参数都包裹进来 这样就可以很方便的处理读写了
```golang
func (r *Server) Register(route string, f HandlerFunc) {
	r.route[route] = withMiddLog(withMiddTime(func(writer http.ResponseWriter, request *http.Request) {
		f(r.createContext(writer, request))
	}))
}
r.Register("/bench", func(c *Context) {
		time.Sleep(time.Second)
		fmt.Println("bench sleep 1 second")
		c.Writer.Write([]byte("hello!\r\n"))
	})
```

###golang框架gin中的实现

####中间件v3.0

>核心理念是将中间件和最后的函数 一视同仁 。通过一个for循环遍历具体的可以参考代码

```golang
func (c *Context) Next() {
	c.index++
	//for中的index++是为了退出循环 否则没法退出
	for ; c.index < len(c.middle); c.index++ {
		c.middle[c.index](c)
	}
}

func withMiddTime() HandleContext {
	return func(c *Context) {
		t := time.Now()
		defer func() {
			fmt.Println("withMiddTime end time", time.Since(t))
		}()
		fmt.Println("withMiddTime start ", time.Since(t))
		c.Next()
	}
}
func (s *Server) Register(path string, f ...HandleContext) {
	handleNew := make([]HandleContext, 0, len(s.handle)+len(f))
	handleNew = append(handleNew, s.handle...)
	handleNew = append(handleNew, f...)
	s.routeHandler(path, func(writer http.ResponseWriter, request *http.Request) {
		s.createContext(writer, request, handleNew).Next()
	})
}
```




