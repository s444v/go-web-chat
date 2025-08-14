package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.GET("/", webHandler)
	//router.Static("/static", "./web")
	router.GET("/api/login", loginHandler)
	router.GET("/api/users", authCookieMiddleware(getUsers))
	router.POST("/api/login", signinHandler)
}

func webHandler(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusFound, "/api/login")
		c.Abort()
		return
	}
	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		c.Redirect(http.StatusFound, "/api/login")
		c.Abort()
		return
	}
	c.File("./web/index.html")
}

func loginHandler(c *gin.Context) {
	c.File("./web/login.html")
}
