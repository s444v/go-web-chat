package api

import "github.com/gin-gonic/gin"

func getUsername(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(200, username)
}
