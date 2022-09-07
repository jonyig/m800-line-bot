package models

type MessageInfo struct {
	UserId  string `json:"user_id" bson:"user_id"`
	Message string `json:"message" bson:"message"`
}
