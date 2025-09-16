package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
