package main

import (
	"fmt"
	"log"
	"time"
	"to-do/config"
	"to-do/controllers"
	"to-do/repository"
	"to-do/routes"
	"to-do/services"

	_ "to-do/docs"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title           Todo API
// @version         1.0
// @description     Todo loyiha API documentation
// @host            localhost:8080
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// 1. ENV o'zgaruvchilarni yukla
	config.LoadConfig()
	fmt.Println("Hello, World!")

	// 2. DB ga ulan
	config.ConnectDB()

	// 3. Repository
	userRepo := repository.NewUserRepository(config.DB)
	categoryRepo := repository.NewCategoryRepository(config.DB)

	// 4. Service
	userService := services.NewUserService(userRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	// 5. Controller
	userController := controllers.NewUserController(userService)
	categoryController := controllers.NewCategoryController(categoryService)

	// 6. Gin router
	r := gin.Default()
	// proxyga ishonma
	r.SetTrustedProxies(nil)

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// swagger ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 7. Routelarni ulash
	routes.SetupRoutes(r, userController, categoryController)

	// 8. Serverni ishga tushur
	port := config.AppConfig.ServerPort
	log.Printf("🚀 Server %s portda ishga tushdi", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server ishga tushmadi: %v", err)
	}
}
