package main

import (
	"context"
	"log"
	"net"

	helloworld2 "github.com/pyhita/golang-base/06_framework/grpc/greeter/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

// server is used to implement .GreeterServer.
type server struct {
	helloworld2.UnimplementedGreeterServer
}

// SayHello implements .GreeterServer.
func (s *server) SayHello(ctx context.Context, in *helloworld2.HelloRequest) (*helloworld2.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld2.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld2.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
