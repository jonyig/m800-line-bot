package config

import (
	"os"
)

type EnvConfig struct {
}

func NewEnvConfig() Config {
	return &EnvConfig{}
}

func (c EnvConfig) GetPort() string {
	return os.Getenv("PORT")
}
func (c EnvConfig) GetMongoDBUsername() string {
	return os.Getenv("MONGO_ROOT_USERNAME")
}
func (c EnvConfig) GetMongoDBPassword() string {
	return os.Getenv("MONGO_ROOT_PASSWORD")
}

func (c EnvConfig) GetLineChannelSecret() string {
	return os.Getenv("CHANNEL_SECRET")
}

func (c EnvConfig) GetLineChannelToken() string {
	return os.Getenv("CHANNEL_TOKEN")
}
