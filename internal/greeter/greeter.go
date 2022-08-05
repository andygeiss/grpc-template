package greeter

import (
	context "context"
)

// Greeter implements the business logic.
type Greeter struct {
	UnimplementedGreeterServer
}

// SayHello greets a user by using his name.
func (a Greeter) SayHello(ctx context.Context, request *SayHelloRequest) (reply *SayHelloReply, err error) {
	return &SayHelloReply{Message: "Hello " + request.Name}, nil
}

// NewGreeter allocates the struct and returns its address.
func NewGreeter() *Greeter {
	return &Greeter{}
}
