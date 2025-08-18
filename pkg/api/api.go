package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.GET("/login", getLoginpage)
	router.POST("/api/login", loginHandler)

	auth := router.Group("/", authCookieMiddleware())
	auth.POST("/api/logout", logoutHandler)
	auth.GET("/api/accounts", getAccounts)
	auth.GET("/mainpage", getMainpage)
	auth.DELETE("api/delete-account", deleteAccount)

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
}
