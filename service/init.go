package service

import (
	greeterPB "github.com/JREAMLU/j-micro/proto"
	microClient "github.com/micro/go-micro/client"
)

const (
	// Greeeter service name
	Greeeter = "go.micro.srv.greeter"
)

var (
	greeterClient greeterPB.GreeterService
)

// InitMicroClient init micro client
func InitMicroClient(c microClient.Client) {
	greeterClient = greeterPB.GreeterServiceClient(Greeeter, c)
}
