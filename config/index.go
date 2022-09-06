package config

import (
	"os"
	"sync"
)

type Config interface {
	GetPort() string
}

var (
	Instance Config
	once     sync.Once
)

func NewConfig() Config {
	var c Config

	if IsNotEnv() {
		c = NewJsonConfig()
	}

	once.Do(func() {
		Instance = c
	})
	return Instance
}

func IsNotEnv() bool {
	return os.Getenv("ENV") == ""
}
