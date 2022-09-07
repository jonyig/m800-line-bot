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
	config      config.Config
	service     *service.LineBotService
	textMapFunc map[string]func() linebot.SendingMessage
}

func NewLineBotHandler(configuration config.Config, service *service.LineBotService) *LineBotHandler {
	return &LineBotHandler{
		config:  configuration,
		service: service,
	}
}

func (h *LineBotHandler) Webhook(context *gin.Context) {
	h.setTextMap()
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
	var result linebot.SendingMessage
	if closure, ok := h.textMapFunc[message]; ok {
		result = closure()
	} else {
		result = h.service.SaveMessage(userID, message)
	}
	if _, err := bot.ReplyMessage(
		replyToken,
		result,
	).Do(); err != nil {
		log.Print(err)
	}
}

func (h *LineBotHandler) setTextMap() {
	h.textMapFunc = map[string]func() linebot.SendingMessage{
		"列出所有訊息": h.service.GetMessagesToBot,
		"廣播訊息":   h.service.BroadcastToBot,
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
func (h *LineBotHandler) Broadcast(c *gin.Context) {
	type Request struct {
		Text string `json:"text" binding:"required"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "text must be pass"})
		return
	}

	code := http.StatusOK

	err := h.service.Broadcast(req.Text)
	if err != nil {
		code = http.StatusInternalServerError
	}
	c.JSON(
		code,
		nil,
	)
}
