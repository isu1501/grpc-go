package main

import (
	"context"
	"fmt"
	"net"

	example "proto/home/djax/Documents/MyFolder/go-grpc"

	"google.golang.org/grpc"
)

type server struct {
	example.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {

	return &example.HelloResponse{
		Message: "Hello, " + req.Name,
	}, nil
}

func main() {
	fmt.Println("Hello.. This is server ")

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Printf("Failed to listen:%v\n", err)
	}

	s := grpc.NewServer()
	example.RegisterExampleServiceServer(s, &server{})
	fmt.Println("Server Started")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}

}
