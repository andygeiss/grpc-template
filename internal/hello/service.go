package hello

import (
	context "context"
)

// Service ...
type Service struct{}

func (a *Service) SayHello(ctx context.Context, request *HelloRequest) (reply *HelloReply, err error) {
	return &HelloReply{Message: "Hello " + request.Name}, nil
}
