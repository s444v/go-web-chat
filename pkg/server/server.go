package server

import (
	"github.com/gin-gonic/gin"
	"github.com/s444v/go-web-chat/pkg/api"
)

func NewServer() *gin.Engine {
	router := gin.Default()
	// router.Use(cors.Default())
	api.HandlersInit(router)
	return router
}
