package hello_test

import (
	"context"
	"testing"

	"github.com/andygeiss/grpc-template/internal/hello"
	"github.com/andygeiss/utils/assert"
)

func Test_Service(t *testing.T) {
	service := hello.NewService()
	reply, err := service.SayHello(context.Background(), &hello.HelloRequest{Name: "Foo"})
	assert.That("error should be nil", t, err, nil)
	assert.That("reply.Message should be Foo", t, reply.Message, "Hello Foo")
}
