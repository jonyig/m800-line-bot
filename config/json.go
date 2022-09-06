package config

import "github.com/spf13/viper"

type JsonConfig struct {
}

func NewJsonConfig() Config {
	return &JsonConfig{}
}

func (c JsonConfig) GetPort() string {
	return viper.GetString("port")
}
