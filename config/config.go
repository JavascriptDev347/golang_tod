package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
}

var AppConfig *Config

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file topilmadi, system env ishlatiladi")
	}

	AppConfig = &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		ServerPort: getEnv("SERVER_PORT"),
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ Majburiy ENV topilmadi: %s", key)
	}
	return value
}
