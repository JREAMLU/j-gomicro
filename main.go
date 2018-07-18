package main

import (
	"time"

	"github.com/JREAMLU/j-micro/controller"
	proto "github.com/JREAMLU/j-micro/proto"
	"github.com/JREAMLU/j-micro/service"
	micro "github.com/micro/go-micro"
)

func main() {
	RunMicroService()
}

// RunMicroService run micro service
func RunMicroService() {
	// Create a new service. Optionally include some options here.
	ms := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.Version("v1"),
	)

	// Init will parse the command line flags.
	ms.Init(
		micro.RegisterTTL(1*time.Second),
		micro.RegisterInterval(1*time.Second),
	)

	// Register handler
	proto.RegisterGreeterHandler(ms.Server(), controller.NewGreeterHandler())

	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
