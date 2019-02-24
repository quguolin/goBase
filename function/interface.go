package main

//博文链接 https://www.flysnow.org/2016/12/30/golang-function-interface.html
import "fmt"

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

//HandlerFunc 类型实现了Do方法
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

type welcome1 string

func (w welcome1) Do(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁。\r\n", w, k, v)
}

type welcome2 string

func (w welcome2) SelfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁。\r\n", w, k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["Janey"] = 10
	persons["Moor"] = 20
	persons["Tango"] = 30

	var w welcome1 = "大家好"
	Each(persons, w)
	var w2 welcome2 = "Hello"
	Each(persons, HandlerFunc(w2.SelfInfo))
	EachFunc(persons, w2.SelfInfo)
}
