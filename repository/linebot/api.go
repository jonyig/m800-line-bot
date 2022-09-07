package repository

import (
	"fmt"
	"m800-line-bot/library"
	"m800-line-bot/models"
)

var (
	LineBotUserInfoAPI  = "https://api.line.me/v2/bot/profile/%s"
	LineBotBroadcastAPI = "https://api.line.me/v2/bot/message/broadcast"
)

type LineBotApiRepository struct {
	httpClient *library.Client
}

func NewLineBotApiRepository(client *library.Client) *LineBotApiRepository {
	return &LineBotApiRepository{
		httpClient: client,
	}
}

func (r *LineBotApiRepository) GetUserInfo(userId string) (username string, err error) {
	type UserInfo struct {
		UserId      string `json:"userId"`
		DisplayName string `json:"displayName"`
	}

	userInfo := &UserInfo{}

	err = r.httpClient.
		SetGetRequest(
			fmt.Sprintf(LineBotUserInfoAPI, userId),
		).
		SetAuthorization().
		Send(userInfo)

	if err != nil {
		return
	}

	username = userInfo.DisplayName
	return
}
func (r *LineBotApiRepository) Broadcast(text string) (err error) {
	message := models.NewBroadcastMessage(text).ToJson()
	err = r.httpClient.
		SetPostRequest(
			LineBotBroadcastAPI,
			message,
		).
		SetAuthorization().
		SetContentTypeJson().
		Send(nil)
	return
}
