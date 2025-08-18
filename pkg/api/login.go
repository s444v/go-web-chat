package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/s444v/go-web-chat/pkg/database"
)

func getLoginPage(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err == nil {
		token, _, err := validateToken(tokenString)
		if err == nil && token.Valid {
			c.Redirect(http.StatusFound, "/mainpage")
		}
	}
	c.File("./web/login.html")
}

func loginHandler(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	ok, roles, err := database.CheckPass(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		Role:     strings.Join(roles, ","),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	c.SetCookie("token", tokenString, 3000, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"roles": roles,
	})
}
