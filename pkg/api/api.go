package api

import (
	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
		//c.File("./web/login.html")
	})
	router.POST("/login", signinHandler)
	auth := router.Group("/", authCookieMiddleware())
	//router.Static("/static", "./web/static")
	auth.GET("/users", getUsers)
}
