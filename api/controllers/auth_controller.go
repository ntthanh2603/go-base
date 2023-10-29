package controllers

import (
	"gin-base/domain"
	"gin-base/services"
	"gin-base/validation"

	"github.com/gin-gonic/gin"
)

func AuthController(r *gin.RouterGroup) {
	authController := r.Group("/auth")
	authService := services.AuthService()
	authController.POST("/login", validation.Body(&domain.LoginDto{}), authService.Login)
}
