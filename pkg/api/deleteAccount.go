package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

func deleteAccount(c *gin.Context) {
	username := c.GetString("username")
	// fmt.Println(username)
	err := database.DeleteAccount(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(
		"token", // имя cookie
		"",      // пустое значение
		-1,      // время жизни (сек) → -1 значит удалить
		"/",     // путь
		"",      // домен ("" = текущий)
		false,   // secure (true только для https)
		true,    // httpOnly
	)
	c.JSON(http.StatusAccepted, gin.H{"messege": "аккаунт удален"})
}
