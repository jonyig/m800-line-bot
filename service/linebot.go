package service

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	repository "m800-line-bot/repository/linebot"
)

type LineBotService struct {
	LineAPI   *repository.LineBotApiRepository
	LineMongo *repository.LineBotMongoRepository
}

func NewLineBotService(lineApi *repository.LineBotApiRepository, lineMongo *repository.LineBotMongoRepository) *LineBotService {
	return &LineBotService{
		LineAPI:   lineApi,
		LineMongo: lineMongo,
	}
}

func (s *LineBotService) SaveMessage(userId string, message string) *linebot.TextMessage {
	username, err := s.LineAPI.GetUserInfo(userId)
	if err != nil {
		log.Print(err)
	}

	s.LineMongo.SaveMessage(username, message)

	return linebot.NewTextMessage(
		message,
	)
}
