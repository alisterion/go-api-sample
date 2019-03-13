package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context)  {
	c.JSON(http.StatusNotFound, gin.H{
		"detail": "Not Found",
	})
}

func NoContent(c *gin.Context)  {
	c.AbortWithStatus(http.StatusNoContent)
}

func Response(c *gin.Context, data interface{}, err error)  {
	if err != nil {
		NotFound(c)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
