package controllers

import (
	"net/http"

	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/models"
	"github.com/gin-gonic/gin"
)

func EditaUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	database.DB.First(&user, id)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
		return
	}

	if err := models.ValidaDados(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
		return
	}

	database.DB.Model(&user).UpdateColumns(user)
	c.JSON(http.StatusOK, user)
}
