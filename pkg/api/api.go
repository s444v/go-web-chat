package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.GET("/", webHandler)
	router.GET("/login", getLoginpage)
	router.POST("/api/login", signinHandler)
	auth := router.Group("/", authCookieMiddleware())
	auth.GET("/api/users", getUsers)
	auth.GET("/mainpage", getMainpage)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
}

func webHandler(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	c.Redirect(http.StatusFound, "/mainpage")
	c.Abort()
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
			return
		}
	}
	c.Redirect(http.StatusFound, "/login")
}
