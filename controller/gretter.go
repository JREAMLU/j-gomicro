package controller

import (
	"context"

	proto "github.com/JREAMLU/jgomicro/proto"
)

// Greeter greeter
type Greeter struct{}

// NewGreeterHandler new greeter
func NewGreeterHandler() *Greeter {
	return &Greeter{}
}

// Hello h w
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
	resp.Greeting = "Hello " + req.Name
	return nil
}
