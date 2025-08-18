package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/database"
)

func addAccount(c *gin.Context) {
	account := &NewAccount{}
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	err := database.AddAccount(account.Username, account.Password, account.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"messege": "аккаунт создан"})
}
