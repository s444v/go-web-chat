package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const LIMIT = 50

type Account struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type NewAccount struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

var jwtKey = []byte("super_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func HandlersInit(router *gin.Engine) {
	router.GET("/login", getLoginPage)
	router.POST("/api/login", loginHandler)
	router.GET("/register", getRegPage)
	router.POST("/api/register", addAccount)
	router.Static("/favicon_io", "./web/favicon_io")

	auth := router.Group("/", authCookieMiddleware())
	auth.POST("/api/logout", logoutHandler)
	auth.GET("/api/accounts", getAccounts)
	auth.GET("/mainpage", getMainPage)
	auth.DELETE("api/delete-account", deleteAccount)
	auth.GET("/ws", wsHandler)

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
}
