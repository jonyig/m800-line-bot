package service

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type LineBotService struct {
}

func NewLineBotService() *LineBotService {
	return &LineBotService{}
}

func (s *LineBotService) SaveMessage(userId string, message string) *linebot.TextMessage {
	log.Print(userId)
	log.Print(message)

	return linebot.NewTextMessage(
		message,
	)
}
