package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "goBase/grpc/v2/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Server represents the gRPC server
type Server struct {
	handlers []grpc.UnaryServerInterceptor
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
	var (
		opt    []grpc.ServerOption
		server *Server
	)
	server = &Server{}
	config := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(10000),
		MaxConnectionAgeGrace: time.Duration(10000),
		Time:                  time.Duration(10000),
		Timeout:               time.Duration(10000),
		MaxConnectionAge:      time.Duration(10000),
	})
	server.handlers = append(server.handlers, server.recovery(), server.time())
	opt = append(opt, config, server.withServerUnaryInterceptor())
	grpcServer := grpc.NewServer(opt...)
	pb.RegisterPingServer(grpcServer, server)
	grpcServer.Serve(lis)
}

func (s *Server) withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(s.interceptor)
}

func (s *Server) recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("recovery")
		resp, err = handler(ctx, req)
		return
	}
}

func (s *Server) time() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("time")
		resp, err = handler(ctx, req)
		return
	}
}

// Authorization unary interceptor function to handle authorize per RPC call
func (s *Server) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var (
		i     int
		chain grpc.UnaryHandler
	)

	n := len(s.handlers)
	if n == 0 {
		return handler(ctx, req)
	}

	chain = func(ic context.Context, ir interface{}) (interface{}, error) {
		if i == n-1 {
			return handler(ic, ir)
		}
		i++
		return s.handlers[i](ic, ir, info, chain)
	}
	return s.handlers[0](ctx, req, info, chain)
}

// authorize function authorizes the token received from Metadata
//func authorize(ctx context.Context) error {
//md, ok := metadata.FromIncomingContext(ctx)
//if !ok {
//	return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
//}
//
//authHeader, ok := md["authorization"]
//if !ok {
//	return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
//}
//
//token := authHeader[0]
//// validateToken function validates the token
//err := fmt.Errorf("validate token error %s",token)
//
//if err != nil {
//	return status.Errorf(codes.Unauthenticated, err.Error())
//}
//return nil
//}
