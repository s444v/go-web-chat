package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const LIMIT = 50

func HandlersInit(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // или конкретные домены, например "http://localhost:3000"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/login", getLoginpage)
	router.POST("/api/login", signinHandler)
	auth := router.Group("/", authCookieMiddleware())
	auth.GET("/api/users", getUsers)
	auth.GET("/mainpage", getMainpage)
	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/mainpage")
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
