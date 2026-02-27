package routes

import (
	"to-do/controllers"
	"to-do/middlewares"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine, todoController *controllers.TodoController) {
	todo := r.Group("/api/todos")
	todo.Use(middlewares.AuthMiddleware())
	{
		todo.POST("", todoController.Create)
		todo.GET("", todoController.FindAll)
		todo.GET("/:id", todoController.FindByID)
		todo.PUT("/:id", todoController.Update)
		todo.DELETE("/:id", todoController.Delete)
	}
}
