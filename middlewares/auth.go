package middlewares

import (
	"net/http"
	"os"
	"strings"
	"to-do/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// Headerdan tokenni olish
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "Token topilmadi")
			c.Abort()
			return
		}

		// 2. "Bearer <token>" dan faqat token qismini olish
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(c, http.StatusUnauthorized, "Token formati noto'g'ri")
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Tokenni verify qilish
		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			utils.Error(c, http.StatusUnauthorized, "Token yaroqsiz yoki muddati o'tgan")
			c.Abort()
			return
		}

		// 4. Claims dan user ma'lumotlarini olish
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.Error(c, http.StatusUnauthorized, "Token ma'lumotlari noto'g'ri")
			c.Abort()
			return
		}

		// 5. Context ga saqlash — keyingi handlerlarda ishlatish uchun
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("role", claims["role"].(string))
		c.Set("email", claims["email"].(string))

		c.Next()
	}
}


