package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"m800-line-bot/routes"
)

func main() {
	r := gin.Default()

	routes.Routes(r)

	err := r.Run(":8000")
	if err != nil {
		return
	}
	log.Print(123)
}
