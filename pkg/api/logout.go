package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func logoutHandler(c *gin.Context) {
	c.SetCookie(
		"token", // имя cookie
		"",      // пустое значение
		-1,      // время жизни (сек) → -1 значит удалить
		"/",     // путь
		"",      // домен ("" = текущий)
		false,   // secure (true только для https)
		true,    // httpOnly
	)
	c.JSON(http.StatusOK, gin.H{"message": "cookie удален"})
}
