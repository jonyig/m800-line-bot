package service

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"m800-line-bot/repository"
)

type LineBotService struct {
	LineAPI *repository.LineApiRepository
}

func NewLineBotService(lineApi *repository.LineApiRepository) *LineBotService {
	return &LineBotService{
		LineAPI: lineApi,
	}
}

func (s *LineBotService) SaveMessage(userId string, message string) *linebot.TextMessage {
	username, err := s.LineAPI.GetUserInfo(userId)
	if err != nil {
		log.Print(err)
	}
	log.Print(username)
	log.Print(message)

	return linebot.NewTextMessage(
		message,
	)
}
