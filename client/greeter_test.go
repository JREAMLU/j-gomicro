package client

import (
	"context"
	"testing"

	proto "github.com/JREAMLU/j-micro/proto"
	microClient "github.com/micro/go-micro/client"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	Convey("Hello", t, func() {
		c := microClient.NewClient()
		greeterClient := NewGreeterClient(c)
		resp, err := greeterClient.Hello("world space")
		if err != nil {
			t.Log("err", err)
			return
		}

		t.Log("resp: ", resp.Greeting)
	})
}

func TestHelloOT(t *testing.T) {
	serviceName := "go.micro.srv.greeter"
	Convey("Hello", t, func() {
		c := microClient.NewClient()
		client := proto.GreeterServiceClient(serviceName, c)
		resp, err := client.Hello(context.Background(), &proto.HelloRequest{
			Name: "LBJ",
		})
		if err != nil {
			t.Log("err", err)
			return
		}

		t.Log("resp: ", resp.Greeting)
	})
}
