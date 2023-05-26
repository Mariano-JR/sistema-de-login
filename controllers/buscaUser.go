package controllers

import (
	"net/http"

	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/models"
	"github.com/gin-gonic/gin"
)

func BuscaUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	database.DB.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found:": "Usuario não encontrado.",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
