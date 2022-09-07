package repository

import (
	"fmt"
	"m800-line-bot/library"
)

var (
	LineUserInfoAPI = "https://api.line.me/v2/bot/profile/%s"
)

type LineApiRepository struct {
	httpClient *library.Client
}

func NewLineApiRepository(client *library.Client) *LineApiRepository {
	return &LineApiRepository{
		httpClient: client,
	}
}

func (r *LineApiRepository) GetUserInfo(userId string) (username string, err error) {
	type UserInfo struct {
		UserId      string `json:"userId"`
		DisplayName string `json:"displayName"`
	}

	userInfo := &UserInfo{}

	err = r.httpClient.
		SetGetRequest(
			fmt.Sprintf(LineUserInfoAPI, userId),
		).
		SetAuthorization().
		Send(userInfo)

	if err != nil {
		return
	}

	username = userInfo.DisplayName
	return
}
