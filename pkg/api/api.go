package api

import (
	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.File("./web/index.html")
	})
	//router.Static("/static", "./web/static")
	router.GET("/users", getUsers)
}
