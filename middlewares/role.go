package middlewares

import (
	"net/http"
	"to-do/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			utils.Error(c, http.StatusForbidden, "Role topilmadi")
			c.Abort()
			return
		}

		// Ruxsat berilgan rolelar ichida bor yo'qligini tekshir
		for _, r := range roles {
			if r == role.(string) {
				c.Next()
				return
			}
		}

		utils.Error(c, http.StatusForbidden, "Bu amalni bajarish uchun ruxsat yo'q")
		c.Abort()
	}
}
