package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.GET("/login", getLoginpage)
	router.POST("/api/login", signinHandler)
	auth := router.Group("/", authCookieMiddleware())
	auth.GET("/api/users", getUsers)
	auth.GET("/mainpage", getMainpage)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
}

func getLoginpage(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err == nil {
		token, err := validateToken(tokenString)
		if err == nil && token.Valid {
			c.Redirect(http.StatusFound, "/mainpage")
			c.Abort()
			return
		}
	}
	c.File("./web/login.html")
}

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
