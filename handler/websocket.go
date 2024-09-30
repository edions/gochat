package handler

import (
	"database/sql"
	"fmt"
	"gochat/model"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan models.Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var db *sql.DB

func InitDB() {
    var err error
	db, err = sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
}

func HandleChat(w http.ResponseWriter, r *http.Request, chatID string) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    Clients[conn] = true

    rows, err := db.Query(`SELECT username, message, timestamp FROM messages WHERE chat_id = ? ORDER BY timestamp ASC`, chatID)
    if err != nil {
        fmt.Println("Error retrieving messages:", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var msg models.Message
        var timestamp string
        err := rows.Scan(&msg.Username, &msg.Message, &timestamp)
        if err != nil {
            fmt.Println("Error scanning message:", err)
            continue
        }

        err = conn.WriteJSON(msg)
        if err != nil {
            fmt.Println("Error sending message:", err)
            continue
        }
    }

    for {
        var msg models.Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            fmt.Println(err)
            delete(Clients, conn)
            return
        }
        msg.ChatID = chatID
        Broadcast <- msg
    }
}


func HandleMessages() {
    for {
        msg := <-Broadcast

        _, err := db.Exec(`INSERT INTO messages (chat_id, username, message) VALUES (?, ?, ?)`, msg.ChatID, msg.Username, msg.Message)
        if err != nil {
            fmt.Println("Error saving message:", err)
        }

        for client := range Clients {
            err := client.WriteJSON(msg)
            if err != nil {
                fmt.Println(err)
                client.Close()
                delete(Clients, client)
            }
        }
    }
}
