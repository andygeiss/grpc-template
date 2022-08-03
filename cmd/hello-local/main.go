package main

import (
	"context"
	"fmt"

	"github.com/andygeiss/grpc-template/internal/hello"
)

func main() {
	service := hello.NewService()
	reply, _ := service.SayHello(context.Background(), &hello.HelloRequest{Name: "Foo"})
	fmt.Println(reply.Message)
}
