package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"m800-line-bot/config"
	"m800-line-bot/service"
	"net/http"
)

type LineBotHandler struct {
	config  config.Config
	service *service.LineBotService
}

func NewLineBotHandler(configuration config.Config, service *service.LineBotService) *LineBotHandler {
	return &LineBotHandler{
		config:  configuration,
		service: service,
	}
}

func (h *LineBotHandler) Webhook(context *gin.Context) {
	bot, err := linebot.New(
		h.config.GetLineChannelSecret(),
		h.config.GetLineChannelToken(),
	)
	if err != nil {
		log.Print(err)
	}

	events, err := bot.ParseRequest(context.Request)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			context.JSON(http.StatusBadRequest, nil)
		}

		context.JSON(http.StatusInternalServerError, nil)
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				h.TextHandler(bot, event.ReplyToken, message)
			}
		}
	}
}

func (h *LineBotHandler) TextHandler(bot *linebot.Client, replyToken string, message *linebot.TextMessage) {
	if _, err := bot.ReplyMessage(
		replyToken,
		h.service.SaveMessage(message.ID, message.Text),
	).Do(); err != nil {
		log.Print(err)
	}
}
