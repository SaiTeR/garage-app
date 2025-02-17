package main

import (
	"garage-app-client-backend/app/controllers"
	"github.com/gin-gonic/gin"
)

// Функция для создания и настройки маршрутов
func SetupRouter() *gin.Engine {
	authController := &controllers.AuthController{}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ХУЙ",
		})
	})

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/logout", authController.Logout)
	}

	//testMiddleware := r.Group("/testmiddleware")
	//testMiddleware.Use(middleware.AuthMiddleware())
	//{
	//	testMiddleware.GET("/auth", func(context *gin.Context) {
	//		context.JSON(200, gin.H{
	//			"message": "Passed auth successfully!",
	//		})
	//	})
	//}

	//taskGroup := r.Group("/task") {
	//	taskGroup.POST("/create", )
	//}

	return r
}
