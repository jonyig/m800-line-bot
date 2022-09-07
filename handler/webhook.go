package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
)

var (
	SingleTermFunc = func(text string) linebot.SendingMessage {
		return linebot.NewTextMessage(
			text,
		)
	}
)

func (h *Handler) Webhook(context *gin.Context) {
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
				TextHandler(bot, event.ReplyToken, message.Text)
			}
		}
	}
}

func TextHandler(bot *linebot.Client, replyToken string, text string) {
	if _, err := bot.ReplyMessage(
		replyToken,
		SingleTermFunc(text),
	).Do(); err != nil {
		log.Print(err)
	}
}
