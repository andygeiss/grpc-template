package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/andygeiss/grpc-template/internal/greeter"
	"google.golang.org/grpc"
)

func main() {
	// connect to the gRPC server
	cc, err := grpc.Dial("127.0.0.1:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	// call the service remotely
	service := greeter.NewGreeterClient(cc)
	reply, _ := service.SayHello(context.Background(), &greeter.SayHelloRequest{Name: "Foo"})
	fmt.Println(reply.Message)
}
