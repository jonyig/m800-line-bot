package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	configFolder "m800-line-bot/config"
	"m800-line-bot/routes"
)

var m800LineBotCmd = &cobra.Command{
	Use: "m800LineBot",

	Run: func(cmd *cobra.Command, args []string) {

		if configFolder.IsNotEnv() {
			err := InitialViper()
			if err != nil {
				return
			}
		}

		config := configFolder.NewConfig()
		r := gin.Default()

		routes.Routes(r)
		err := r.Run(fmt.Sprintf(":%s", config.GetPort()))
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
