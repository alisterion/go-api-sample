package controllers

import (
	"github.com/gin-gonic/gin"
	"sample/pkg/api"
	"sample/pkg/models"
)

func ListUser(c *gin.Context)  {
	var users []models.User
	err := models.GetAll(&users)
	api.Response(c, users, err)
}

func CreateUser(c *gin.Context)  {
	var user models.User
	_ = c.BindJSON(&user)
	err := models.Create(&user)
	api.Response(c, user, err)
}

func RetrieveUser(c *gin.Context)  {
	var user models.User
	err := models.GetOne(&user, c.Param("id"))
	api.Response(c, user, err)
}

func UpdateUser(c *gin.Context)  {
	var user models.User

	if models.GetOne(&user, c.Param("id")) != nil {
		api.NotFound(c); return
	}
	_ = c.BindJSON(&user)

	err := models.Save(&user)
	api.Response(c, user, err)
}

func DeleteUser(c *gin.Context)  {
	var user models.User
	id := c.Param("id")

	if models.GetOne(&user, id) != nil {
		api.NotFound(c); return
	}

	_ = models.Delete(&user, id)
	api.NoContent(c)
}
