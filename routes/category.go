package routes

import (
	"to-do/controllers"
	"to-do/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, categoryController *controllers.CategoryController) {
	category := r.Group("/api/categories")

	{
		// Hammaga ochiq
		category.GET("", categoryController.FindAll)
		category.GET("/:id", categoryController.FindById)

		// Faqat admin uchunlari
		admin := category.Group("")
		admin.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("admin"))
		{
			admin.POST("", categoryController.Create)
			admin.PUT("/:id", categoryController.Update)
			admin.DELETE("/:id", categoryController.Delete)
		}
	}
}
