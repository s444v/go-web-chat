package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

func getAccounts(c *gin.Context) {
	userId, err := database.GetUserId(c.GetString("username"))
	if err != nil {
		return
	}
	accounts, err := database.GetAccounts(userId, LIMIT, c.Query("search"))
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, accounts)
}
