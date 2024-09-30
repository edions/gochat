package models

type Message struct {
	ChatID string `json:"chatid"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
