package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(route *gin.Engine) {
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})
}
