package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"m800-line-bot/config"
	"m800-line-bot/models"
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
				h.TextHandler(
					bot,
					event.ReplyToken,
					event.Source.UserID,
					message.Text,
				)
			}
		}
	}
}

func (h *LineBotHandler) TextHandler(bot *linebot.Client, replyToken string, userID string, message string) {
	if _, err := bot.ReplyMessage(
		replyToken,
		h.service.SaveMessage(userID, message),
	).Do(); err != nil {
		log.Print(err)
	}
}

func (h *LineBotHandler) GetMessage(c *gin.Context) {
	var response models.Response
	code := http.StatusOK

	list, err := h.service.GetMessages()
	if err != nil {
		code = http.StatusInternalServerError
	}

	response.SetData(list).SetErr(err)

	c.JSON(
		code,
		response,
	)
}
