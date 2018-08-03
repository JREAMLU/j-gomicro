package main

import (
	"github.com/JREAMLU/j-micro/config"
	"github.com/JREAMLU/j-micro/controller"
	proto "github.com/JREAMLU/j-micro/proto"
	"github.com/JREAMLU/j-micro/service"

	"github.com/JREAMLU/j-kit/go-micro/util"
)

func main() {
	// load config
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	RunMicroService(conf)
}

// RunMicroService run micro service
func RunMicroService(conf *config.GreeterConfig) {
	ms := util.NewMicroService(conf.Config)

	// Register handler
	proto.RegisterGreeterHandler(ms.Server(), controller.NewGreeterHandler())

	// init client
	service.InitMicroClient(ms.Client())

	// Run the server
	if err := ms.Run(); err != nil {
		panic(err)
	}
}
