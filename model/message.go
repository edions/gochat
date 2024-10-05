package models

type Message struct {
	ChatID string `json:"chatid"`
	UserID string `json:"userid"`
	Message  string `json:"message"`
	File string `json:"file"`
}
