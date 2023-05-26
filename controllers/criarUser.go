package controllers

import (
	"net/http"

	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/models"
	"github.com/gin-gonic/gin"
)

func CriarNovoUser(c *gin.Context) {
	var user models.User
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

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
