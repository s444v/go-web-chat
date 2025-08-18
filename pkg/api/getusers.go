package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

type Account struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

func getAccounts(c *gin.Context) {
	accounts, err := database.GetAccounts(LIMIT, c.Query("search"))
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, accounts)
}
