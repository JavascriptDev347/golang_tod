package main

import (
	"fmt"
	"log"
	"to-do/config"
	"to-do/controllers"
	"to-do/repository"
	"to-do/routes"
	"to-do/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. ENV o'zgaruvchilarni yukla
	config.LoadConfig()
	fmt.Println("Hello, World!")

	// 2. DB ga ulan
	config.ConnectDB()

	// 3. Repository
	userRepo := repository.NewUserRepository(config.DB)

	// 4. Service
	userService := services.NewUserService(userRepo)

	// 5. Controller
	userController := controllers.NewUserController(userService)

	// 6. Gin router
	r := gin.Default()
	// proxyga ishonma
	r.SetTrustedProxies(nil)

	// 7. Routelarni ulash
	routes.SetupRoutes(r, userController)

	// 8. Serverni ishga tushur
	port := config.AppConfig.ServerPort
	log.Printf("🚀 Server %s portda ishga tushdi", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server ishga tushmadi: %v", err)
	}
}
