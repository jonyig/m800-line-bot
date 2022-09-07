package handler

import "m800-line-bot/config"

type Handler struct {
	config config.Config
}

func NewHandler(configuration config.Config) *Handler {
	return &Handler{
		config: configuration,
	}
}
