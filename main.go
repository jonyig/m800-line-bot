/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"m800-line-bot/cmd"
	"m800-line-bot/config"
	"m800-line-bot/handler"
	"m800-line-bot/library"
	repository "m800-line-bot/repository/linebot"
	"m800-line-bot/routes"
	"m800-line-bot/service"
	"m800-line-bot/storage"
)

func main() {
	if config.IsNotEnv() {
		err := cmd.InitialViper()
		if err != nil {
			fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
			return
		}
	}

	mongoClient := library.GetMongoDBInstance()
	storage.SetMessage(mongoClient)

	configuration := config.NewConfig()
	client := library.NewClient(configuration.GetLineChannelToken())
	lineApi := repository.NewLineBotApiRepository(client)
	lineMongo := repository.NewLineBotMongoRepository(mongoClient)
	LineBotService := service.NewLineBotService(
		lineApi,
		lineMongo,
	)
	h := handler.NewLineBotHandler(
		configuration,
		LineBotService,
	)

	r := gin.Default()

	routes.Routes(r, h)
	err := r.Run(fmt.Sprintf(":%s", configuration.GetPort()))
	if err != nil {
		return
	}
}
