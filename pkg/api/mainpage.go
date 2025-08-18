package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMainpage(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err == nil {
		token, err := validateToken(tokenString)
		if err == nil && token.Valid {
			c.File("./web/index.html")
			c.Abort()
			return
		}
	}
	c.Redirect(http.StatusFound, "/login")
}
