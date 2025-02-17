package middleware

import (
	"garage-app-client-backend/app/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware проверяет наличие и действительность токена в Redis
func AuthMiddleware() gin.HandlerFunc {
	jwtService := services.NewJWTService()
	redisService := services.NewRedisService()

	return func(c *gin.Context) {
		// Извлекаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		// Токен должен быть в формате "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token format",
			})
			c.Abort()
			return
		}

		token := tokenParts[1]

		// Проверяем наличие токена в Redis
		tokenAmount, err := redisService.CheckToken(c, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error checking token in Redis",
			})
			c.Abort()
			return
		}
		log.Println(tokenAmount)
		if tokenAmount == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Декодируем токен и проверяем его
		_, err = jwtService.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		// Если все проверки прошли успешно, продолжаем выполнение запроса
		c.Next()
	}
}
