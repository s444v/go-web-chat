package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

func getAccounts(c *gin.Context) {
	accounts, err := database.GetAccounts(LIMIT, c.Query("search"))
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, accounts)
}
