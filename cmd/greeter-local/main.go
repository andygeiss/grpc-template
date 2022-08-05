package main

import (
	"context"
	"fmt"

	"github.com/andygeiss/grpc-template/internal/greeter"
)

func main() {
	// call the service locally
	service := greeter.NewGreeter()
	reply, _ := service.SayHello(context.Background(), &greeter.SayHelloRequest{Name: "Foo"})
	fmt.Println(reply.Message)
}
