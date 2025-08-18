package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRegPage(c *gin.Context) {
	c.Status(http.StatusOK)
	c.File("./web/registr.html")
}
