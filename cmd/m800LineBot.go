package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"m800-line-bot/config"
	"m800-line-bot/handler"
	"m800-line-bot/library"
	"m800-line-bot/repository/linebot"
	"m800-line-bot/routes"
	"m800-line-bot/service"
	"m800-line-bot/storage"
)

var m800LineBotCmd = &cobra.Command{
	Use: "m800LineBot",

	Run: func(cmd *cobra.Command, args []string) {

		if config.IsNotEnv() {
			err := InitialViper()
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
		service := service.NewLineBotService(
			lineApi,
			lineMongo,
		)
		h := handler.NewLineBotHandler(
			configuration,
			service,
		)

		r := gin.Default()

		routes.Routes(r, h)
		err := r.Run(fmt.Sprintf(":%s", configuration.GetPort()))
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(m800LineBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// m800LineBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// m800LineBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
