package main

import (
	"context"
	"fmt"
	"log"

	"github.com/andygeiss/grpc-template/internal/hello"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	service := hello.NewGreeterClient(cc)
	reply, _ := service.SayHello(context.Background(), &hello.HelloRequest{Name: "Foo"})
	fmt.Println(reply.Message)
}
