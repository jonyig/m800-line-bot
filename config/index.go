package config

import (
	"os"
)

type Config interface {
	GetPort() string
	GetMongoDBUsername() string
	GetMongoDBPassword() string
	GetLineChannelSecret() string
	GetLineChannelToken() string
}

func NewConfig() Config {
	if IsNotEnv() {
		return NewJsonConfig()
	}

	return NewEnvConfig()
}

func IsNotEnv() bool {
	return os.Getenv("ENV") == ""
}
