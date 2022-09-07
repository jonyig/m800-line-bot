package service

import (
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"m800-line-bot/models"
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

func (s *LineBotService) GetMessagesToBot() linebot.SendingMessage {
	list, err := s.LineMongo.GetMessages()
	if err != nil {
		return linebot.NewTextMessage("service fail")
	}
	jsonString, _ := json.Marshal(list)
	return linebot.NewTextMessage(string(jsonString))
}

func (s *LineBotService) GetMessages() (list []models.MessageInfo, err error) {
	return s.LineMongo.GetMessages()
}

func (s *LineBotService) BroadcastToBot() linebot.SendingMessage {
	err := s.LineAPI.Broadcast("Hello everyone")
	if err != nil {
		return linebot.NewTextMessage("service fail")
	}
	return linebot.NewTextMessage("")
}

func (s *LineBotService) Broadcast(text string) error {
	return s.LineAPI.Broadcast(text)
}
