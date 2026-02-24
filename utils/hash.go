package utils

import (
	"errors"
	"os"
	"time"
	"to-do/dto"
	"to-do/models"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("token yaratishda xatolik")
	}

	return signed, nil
}

func BuildAuthResponse(token string, user *models.User) *dto.AuthResponse {
	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  string(user.Role),
		},
	}
}
