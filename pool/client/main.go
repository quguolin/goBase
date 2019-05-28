package main

import (
	"fmt"
	"log"
	"net"
	"sync/atomic"
	"time"
)

func connHandler(c net.Conn) {
	buf := make([]byte,1024)
	for {
		cnt, err := c.Read(buf)
		if err != nil{
			log.Println(err)
			atomic.AddInt64(&total,-1)
			return
		}
		fmt.Println(string(buf[:cnt]))
	}

}

var total int64

func count()  {
	for{
		fmt.Println("count is ",total)
		time.Sleep(time.Second)
	}
}
func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil{
		panic(err)
	}

	go count()

	for {
		conn, err := server.Accept()
		if err != nil{
			log.Fatal(err)
			return
		}
		atomic.AddInt64(&total, 1)
		go connHandler(conn)
	}
}
