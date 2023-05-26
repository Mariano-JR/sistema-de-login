package controllers

import (
	"net/http"

	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/models"
	"github.com/gin-gonic/gin"
)

func DeletaUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	database.DB.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{
		"data:": "Usuario deletado com sucesso.",
	})
}
