package routes

import (
	"github.com/gin-gonic/gin"
	"m800-line-bot/handler"
	"net/http"
)

func Routes(route *gin.Engine, h *handler.LineBotHandler) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})
	route.POST("/webhook", h.Webhook)
	route.GET("/messages", h.GetMessage)
	route.POST("/broadcast", h.Broadcast)
}
