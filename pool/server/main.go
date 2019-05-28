package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var(
	connChans chan net.Conn
	mu sync.Mutex
)

func newConn() (net.Conn,error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080") // TCP连接
	if err != nil{
		return nil,err
	}
	return conn,err
}

func getCon() net.Conn {
	mu.Lock()
	conn := <-connChans
	mu.Unlock()
	return conn
}

func putConn(conn net.Conn)  {
	mu.Lock()
	connChans<-conn
	mu.Unlock()
}

func get()  {
	fmt.Println("start get connect")
	c := getCon()
	c.Write([]byte("hello"))
	fmt.Println("get connect")
	time.Sleep(10*time.Second)
	putConn(c)
	fmt.Println("put connect")
}

func new()  {
	for i:=0 ;i<10;i++{
		c,err := newConn()
		if err != nil{
			fmt.Println(err)
			continue
		}
		connChans<-c
	}
}

func count()  {
	for{
		fmt.Println("len is ",len(connChans),"cap is ",cap(connChans))
		time.Sleep(time.Second)
	}
}

func main()  {
	connChans = make(chan net.Conn,10)
	new()
	go count()
	go get()
	for{
		time.Sleep(time.Second)
	}
}
