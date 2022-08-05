package greeter_test

import (
	"context"
	"testing"

	"github.com/andygeiss/grpc-template/internal/greeter"
	"github.com/andygeiss/utils/assert"
)

func Test_Service(t *testing.T) {
	g := greeter.NewGreeter()
	reply, err := g.SayHello(context.Background(), &greeter.SayHelloRequest{Name: "Foo"})
	assert.That("error should be nil", t, err, nil)
	assert.That("reply.Message should be Foo", t, reply.Message, "Hello Foo")
}
