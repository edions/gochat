package api

import (
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/ws", HandleConnections)

	go HandleMessages()
}
