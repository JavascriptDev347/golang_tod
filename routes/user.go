package routes

import (
	"to-do/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userController *controllers.UserController) {
	auth := r.Group("/api/auth")

	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
	}
}
