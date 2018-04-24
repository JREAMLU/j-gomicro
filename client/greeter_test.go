package client

import (
	"fmt"
	"log"
	"testing"

	microClient "github.com/micro/go-micro/client"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHello(t *testing.T) {
	Convey("Hello", t, func() {
		c := microClient.NewClient()
		greeterClient := NewGreeterClient(c)
		resp, err := greeterClient.Hello("world space")
		if err != nil {
			log.Println("err", err)
			return
		}

		fmt.Println("++++++++++++: ", resp.Greeting)
	})
}
