package controllers

import (
	"net/http"

	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/models"
	"github.com/gin-gonic/gin"
)

func CriarNovoPost(c *gin.Context) {
	var (
		user models.User
		post models.Post
	)

	id := c.Param("id")

	database.DB.First(&user, id)
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
		return
	}

	database.DB.Create(&post)

}
