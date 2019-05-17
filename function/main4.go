package main

import "fmt"

type S struct {
}

func (s *S) Service(name string) {
	fmt.Println("Service ", name)
}

func newService() *S {
	return &S{}
}

func withService(fn func(s *S)) func() {
	return func() {
		fn(newService())
	}
}

func main() {
	withService(func(s *S) {
		s.Service("first")
	})()
	withService(func(s *S) {
		s.Service("second")
	})()
	withService(func(s *S) {
		s.Service("third")
	})()
	//newService().Service()
	//serviceFirst()
	//serviceSecond()
	//serviceThird()
}

//func serviceFirst() {
//	newService().Service("service First")
//}
//
//func serviceSecond() {
//	newService().Service("service Second")
//}
//
//func serviceThird() {
//	newService().Service("service Third")
//}
