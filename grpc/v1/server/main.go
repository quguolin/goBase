package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "goBase/grpc/v1/api"

	"google.golang.org/grpc"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &pb.PingMessage{Greeting: "你好！世界！"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPingServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
