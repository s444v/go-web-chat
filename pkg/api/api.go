package api

import (
	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
	})
	router.POST("/login", signinHandler)
	auth := router.Group("/api", authCookieMiddleware())
	//router.Static("/static", "./web/static")
	auth.GET("/users", getUsers)
}
