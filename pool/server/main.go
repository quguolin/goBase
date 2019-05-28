package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type ConnPool struct {
	mu sync.Mutex
	connChans chan net.Conn
}

func newConn() *ConnPool {
	c := &ConnPool{
		connChans:make(chan net.Conn,10),
	}
	for i:=0 ;i<10;i++{
		conn,err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil{
			fmt.Println(err)
			continue
		}
		c.connChans<-conn
	}
	return c
}

func (c *ConnPool)getCon() net.Conn {
	c.mu.Lock()
	conn := <-c.connChans
	c.mu.Unlock()
	return conn
}

func (c *ConnPool)putConn(conn net.Conn)  {
	c.mu.Lock()
	c.connChans<-conn
	c.mu.Unlock()
}

func (c *ConnPool)get()  {
	fmt.Println("start get connect")
	conn := c.getCon()
	conn.Write([]byte("hello"))
	fmt.Println("get one connect")
	time.Sleep(10*time.Second)
	c.putConn(conn)
	fmt.Println("put connect")
}


func count(c *ConnPool)  {
	for{
		fmt.Println("len is ",len(c.connChans),"cap is ",cap(c.connChans))
		time.Sleep(time.Second)
	}
}

func main()  {
	c := newConn()
	go count(c)
	go c.get()
	for{
		time.Sleep(time.Second)
	}
}
