package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func getUsers(c *gin.Context) {
	users, err := database.GetUsers(LIMIT, c.Query("search"))
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}
