package controllers

import (
	"fmt"
	"garage-app-client-backend/app/repositories"
	"garage-app-client-backend/app/requests"
	"garage-app-client-backend/app/services"
	"garage-app-client-backend/app/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AuthController struct{}

// Регистрация пользователя
func (ctrl *AuthController) Register(c *gin.Context) {
	var request requests.RegisterRequest
	userRepository := repositories.NewUserRepository()
	jwtService := services.NewJWTService()
	redisService := services.NewRedisService()

	validators.Validate(c, &request)
	if c.IsAborted() {
		return
	}

	userID, err := userRepository.CreateUser(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add user to DB",
		})
		return
	}

	// Генерируем токен
	token, err := jwtService.GenerateToken(strconv.FormatUint(uint64(userID), 10))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate JWT token",
		})
		return
	}

	// Сохраняем токен в Redis
	err = redisService.StoreToken(fmt.Sprintf("%d", userID), token, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to store JWT token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Client registered successfully!",
		"JWTtoken": token,
	})
}

func (ctrl *AuthController) Login(c *gin.Context)  {}
func (ctrl *AuthController) Logout(c *gin.Context) {}

//func (ctrl *AuthController) Refresh(c *gin.Context) {}
