package main

import (
	"fmt"
	"gochat/handler"
	"net/http"

	"github.com/google/uuid"
)

func InitializeRoutes() {
    http.HandleFunc("GET /", HomePage)
    http.HandleFunc("GET /ws", HandleConnections)
    http.HandleFunc("GET /ws/", HandleChatHistory)
    http.HandleFunc("POST /upload", handler.HandlerUpload)

    fs := http.FileServer(http.Dir(uploadPath))
    http.Handle("GET /uploads/", http.StripPrefix("/uploads/", fs))

    go handler.HandleMessages()
}

const uploadPath = "./uploads"

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
