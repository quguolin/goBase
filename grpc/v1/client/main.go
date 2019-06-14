package main

import (
	"context"
	"fmt"
	"time"

	pb "goBase/grpc/v1/api"

	"google.golang.org/grpc"
)
func main()  {
	conn, err := grpc.Dial("127.0.0.1:9001",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewPingClient(conn)
	for{
		msg,err := client.SayHello(context.Background(),&pb.PingMessage{
			Greeting:"hello world!",
		})
		if err != nil{
			panic(err)
		}
		fmt.Println(msg.Greeting)
		time.Sleep(time.Second)
	}
}
