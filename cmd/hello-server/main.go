package main

import (
	"fmt"
	"log"
	"net"

	"github.com/andygeiss/grpc-template/internal/hello"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	service := hello.NewService()
	hello.RegisterGreeterServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
