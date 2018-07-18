package client

import (
	"context"

	proto "github.com/JREAMLU/j-micro/proto"
	"github.com/micro/go-micro/client"
)

// ServiceName servcie
const ServiceName = "go.micro.srv.greeter"

// GreeterClient client
type GreeterClient struct {
	greeterClient proto.GreeterService
}

// NewGreeterClient new client
func NewGreeterClient(c client.Client) *GreeterClient {
	return &GreeterClient{
		greeterClient: proto.GreeterServiceClient(ServiceName, c),
	}
}

// Hello hello
func (g *GreeterClient) Hello(name string) (resp *proto.HelloResponse, err error) {
	return g.greeterClient.Hello(context.Background(), &proto.HelloRequest{
		Name: name,
	})
}
