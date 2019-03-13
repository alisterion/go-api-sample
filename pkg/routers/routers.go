package routers

import (
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		SetupUser(v1.Group("/users"))
	}
	return r
}
