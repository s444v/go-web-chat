package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

func messagesHandler(c *gin.Context) {
	chatID, _ := strconv.Atoi(c.Query("chat_id"))

	messages, err := database.GetMessages(chatID, LIMIT)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, messages)
}
