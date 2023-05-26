package routes

import (
	"github.com/Mariano-JR/sistema-de-login/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/users", controllers.Users)
	r.GET("/users/:id", controllers.BuscaUser)
	r.POST("/users", controllers.CriarNovoUser)
	r.DELETE("/users/:id", controllers.DeletaUser)
	r.PATCH("/users/:id", controllers.EditaUser)

	r.POST("/users/:id/post", controllers.CriarNovoPost)

	r.Run()
}
