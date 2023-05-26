package controllers

import (
	"net/http"

	"github.com/Mariano-JR/sistema-de-login/database"
	"github.com/Mariano-JR/sistema-de-login/models"
	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
