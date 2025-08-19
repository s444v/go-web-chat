package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var clients = make(map[string]*Client)

func wsHandler(c *gin.Context) {
	userID := c.Query("id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("upgrade err:", err)
		return
	}
	client := &Client{ID: userID, Conn: conn}
	clients[userID] = client
	defer func() {
		conn.Close()
		delete(clients, userID)
	}()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read error", err)
			break
		}
		fmt.Printf("messege from %s: %s\n", userID, msg)
	}
}
