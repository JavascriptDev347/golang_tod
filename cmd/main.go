package main

import (
	"fmt"
	"log"
	"to-do/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. ENV o'zgaruvchilarni yukla
	config.LoadConfig()
	fmt.Println("Hello, World!")

	// 2. DB ga ulan
	config.ConnectDB()

	// 3. Gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 4. Serverni ishga tushur
	port := config.AppConfig.ServerPort
	log.Printf("🚀 Server %s portda ishga tushdi", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server ishga tushmadi: %v", err)
	}
}
