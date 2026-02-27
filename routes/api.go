package routes

import (
	"to-do/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine,
	userController *controllers.UserController,
	categoryController *controllers.CategoryController,
	todoController *controllers.TodoController,

) {
	UserRoutes(r, userController)
	CategoryRoutes(r, categoryController)
	TodoRoutes(r, todoController)

}
