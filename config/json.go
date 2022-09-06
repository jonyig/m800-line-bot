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
func (c JsonConfig) GetMongoDBUsername() string {
	return viper.GetString("mongo.username")
}
func (c JsonConfig) GetMongoDBPassword() string {
	return viper.GetString("mongo.password")
}
