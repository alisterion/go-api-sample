package routers

import (
	"github.com/gin-gonic/gin"
	"sample/pkg/controllers"
)

func SetupUser(r *gin.RouterGroup)  {
	r.GET("/", controllers.ListUser)
	r.POST("/", controllers.CreateUser)
	r.GET("/:id", controllers.RetrieveUser)
	r.PUT("/:id", controllers.UpdateUser)
	r.DELETE("/:id", controllers.DeleteUser)
}