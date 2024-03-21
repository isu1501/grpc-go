package main

import (
	"context"
	"fmt"

	example "proto/home/djax/Documents/MyFolder/go-grpc"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello.. This is client")

	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect", err)
	}
	defer conn.Close()

	client := example.NewExampleServiceClient(conn)

	resp, err := client.SayHello(context.Background(), &example.HelloRequest{Name: "Max"})
	if err != nil {
		fmt.Println("Failed to call SayHello", err)
	}
	fmt.Printf("Response from server %s\n", resp.Message)

}
