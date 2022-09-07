package models

import "encoding/json"

type BroadcastMessage struct {
	Message []BroadcastDetail `json:"messages"`
}

type BroadcastDetail struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func NewBroadcastMessage(text string) *BroadcastMessage {
	return &BroadcastMessage{
		Message: []BroadcastDetail{
			{
				Type: "text",
				Text: text,
			},
		},
	}
}

func (m *BroadcastMessage) ToJson() (result []byte) {
	result, _ = json.Marshal(m)
	return
}
