package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.GET("/", authCookieMiddleware(webHandler))
	router.Static("/static", "./web")
	router.GET("/api/users", authCookieMiddleware(getUsers))
	router.POST("/api/login", signinHandler)
}

func webHandler(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Redirect(http.StatusFound, "/static/login.html")
		c.Abort()
		return
	}
	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		// Обычная страница — редирект на login.html
		c.Redirect(http.StatusFound, "/login.html")
		c.Abort()
		return
	}

}
