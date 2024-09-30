package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
    "gochat/handler"
)

func InitializeRoutes() {
    http.HandleFunc("GET /", HomePage)
    http.HandleFunc("GET /ws", HandleConnections)
    http.HandleFunc("GET /ws/", HandleChatHistory)
    go handler.HandleMessages()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Go Chat!")
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
    chatID := uuid.New().String()
    handler.HandleChat(w, r, chatID)
}

func HandleChatHistory(w http.ResponseWriter, r *http.Request) {
    chatID := r.URL.Path[len("/ws/"):]
    handler.HandleChat(w, r, chatID)
}
