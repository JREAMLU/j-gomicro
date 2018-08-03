package config

import (
	"github.com/JREAMLU/j-kit/go-micro/util"
)

const (
	name    = "greeter"
	version = "v1"
)

// GreeterConfig greeter config
type GreeterConfig struct {
	*util.Config

	// S1 custom define config
	Greeter struct {
		Secret string
	}
}

// Load load config
func Load() (*GreeterConfig, error) {
	// load redis mysql elastic client

	// load parent config
	greeterConfig := &GreeterConfig{}
	err := util.LoadCustomConfig("10.200.202.35:8500", name, version, greeterConfig)

	return greeterConfig, err
}
