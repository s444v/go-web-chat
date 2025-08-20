package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/s444v/go-web-chat/pkg/database"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	UserId int
	Conn   *websocket.Conn
}

type Message struct {
	ChatID int    `json:"chat_id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

var clients = make(map[int]*Client)

func wsHandler(c *gin.Context) {
	username := c.GetString("username")
	userId, err := database.GetUserId(username)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade err:", err)
		return
	}
	client := &Client{UserId: userId, Conn: conn}
	clients[userId] = client
	defer func() {
		conn.Close()
		delete(clients, userId)
	}()
	for {
		var msg Message
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("cant save message", err)
			break
		}

		// сохраняем в БД
		err = database.AddMessage(msg.ChatID, userId, msg.Text)
		if err != nil {
			fmt.Println("cant save message", err)
			break
		}

		// проставляем отправителя перед рассылкой
		msg.Sender = username

		fmt.Printf("message from %s: %s\n", username, msg.Text)

		// рассылаем всем клиентам
		for _, client := range clients {
			if err := client.Conn.WriteJSON(msg); err != nil {
				fmt.Println("cant send message", err)
			}
		}
	}
}
