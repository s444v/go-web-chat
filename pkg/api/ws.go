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
	Username string
	Conn     *websocket.Conn
}

type Message struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

var clients = make(map[string]*Client)

func wsHandler(c *gin.Context) {
	username := c.GetString("username")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade err:", err)
		return
	}
	client := &Client{Username: username, Conn: conn}
	clients[username] = client
	defer func() {
		conn.Close()
		delete(clients, username)
	}()
	for {
		var msg Message
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("cant save messege", err)
			continue
		}
		err = database.AddMessege(username, msg.To, msg.Text)
		if err != nil {
			fmt.Println("cant save messege", err)
			continue
		}
		fmt.Printf("messege from %s: %s\n", username, msg)
	}
}
